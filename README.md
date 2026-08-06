# Arquitetura de composição: proxies de protocolo, contexto interno e `WindowState` agnóstico

A arquitetura proposta separa claramente as responsabilidades em três camadas distintas:

```text
Aplicação
    │
    ▼
Window (API pública)
    │
    ▼
WindowDriver (comandos)
    │
    ▼
WindowContext (composição interna)
    │
    ├───────────────┐
    ▼               ▼
Proxies Wayland   WindowState
```

O princípio central é manter os **proxies do protocolo** completamente desacoplados da lógica da biblioteca. Eles não conhecem a aplicação, não conhecem o `WindowState` e não sabem qual entidade de alto nível representam. Sua única responsabilidade é traduzir a comunicação com o protocolo Wayland:

* converter chamadas da biblioteca em mensagens do protocolo;
* converter mensagens recebidas do protocolo em eventos internos.

Toda a lógica de composição fica concentrada no `WindowContext`.

---

# 1. Proxies de protocolo

Os proxies (`wl_surface`, `xdg_surface`, `xdg_toplevel`, etc.) representam diretamente os objetos definidos pelo protocolo Wayland.

Eles conhecem apenas detalhes de comunicação:

* `object_id`;
* `opcode`;
* serialização;
* desserialização;
* envio de mensagens;
* tradução de eventos.

Eles **não possuem estado da janela nem lógica de composição**.

Exemplo:

```go
type XdgToplevel struct {
    objectId uint32
    send     func(*protocol.Message) error
}
```

Quando recebem um evento como:

```text
xdg_toplevel.configure
```

o proxy não modifica nenhuma estrutura de estado. Ele apenas converte a mensagem do protocolo em um evento interno:

```text
Mensagem Wayland
        │
        ▼
ToplevelConfigureEvent
```

Exemplo:

```go
type ToplevelConfigureEvent struct {
    Width  uint32
    Height uint32
}
```

Dessa forma, os proxies permanecem simples, reutilizáveis e independentes da arquitetura da biblioteca.

---

# 2. Dispatcher como roteador

O `Dispatcher` é responsável por descobrir qual entidade da aplicação é proprietária de cada objeto do protocolo.

Uma aplicação pode possuir várias janelas, e cada janela é composta por diversos objetos Wayland.

```text
Window A
    ├── wl_surface    (id 10)
    ├── xdg_surface   (id 11)
    └── xdg_toplevel  (id 12)

Window B
    ├── wl_surface    (id 20)
    ├── xdg_surface   (id 21)
    └── xdg_toplevel  (id 22)
```

Por isso, o dispatcher não pode encaminhar todos os eventos para um único `WindowState`.

Em vez disso, ele mantém uma associação entre cada `object_id` e seu respectivo contexto:

```go
map[uint32]EventTarget
```

Exemplo:

```text
10 → WindowContext A
11 → WindowContext A
12 → WindowContext A

20 → WindowContext B
21 → WindowContext B
22 → WindowContext B
```

Quando chega uma mensagem como:

```text
object_id = 12
opcode    = configure
```

o fluxo é:

```text
Dispatcher
    │
    ▼
Proxy XdgToplevel
    │
    ▼
ToplevelConfigureEvent
    │
    ▼
WindowContext A
```

O dispatcher atua apenas como um roteador entre objetos do protocolo e o contexto responsável por eles.

---

# 3. `WindowContext` como proxy interno

O `WindowContext` é a camada de composição da biblioteca.

Diferentemente dos proxies Wayland, ele não representa um objeto do protocolo, mas sim uma **entidade lógica** formada pela combinação de vários objetos.

Uma janela Wayland não existe como um único objeto:

```text
wl_surface
      +
xdg_surface
      +
xdg_toplevel
```

O `WindowContext` reúne todos esses componentes em uma única representação coerente:

```go
type WindowContext struct {
    State        *WindowState

    WlSurface    *proxies.WlSurface
    XdgSurface   *proxies.XdgSurface
    XdgToplevel  *proxies.XdgToplevel
}
```

Esse contexto passa a ser o proprietário da composição completa da janela.

---

# 4. `WindowState` como estado agnóstico

O `WindowState` representa exclusivamente o estado lógico de uma janela dentro da biblioteca.

Exemplo:

```go
type WindowState struct {
    width   uint32
    height  uint32
    visible bool
    title   string
}
```

Ele não conhece nenhum detalhe do protocolo Wayland.

Portanto, não possui referências para:

* `wl_surface`;
* `xdg_toplevel`;
* `object_id`;
* `serial`;
* buffers;
* mensagens do protocolo.

Sua única responsabilidade é responder à pergunta:

> Qual é o estado atual desta janela?

Essa separação torna o estado completamente independente da implementação do backend gráfico.

---

# 5. Fluxo completo de eventos

Considere o envio da seguinte mensagem pelo compositor:

```text
xdg_toplevel.configure(800, 600)
```

O fluxo completo é:

```text
Compositor Wayland
        │
        ▼
Dispatcher
        │
        ▼
Proxy XdgToplevel
        │
        ▼
ToplevelConfigureEvent
        │
        ▼
WindowContext
        │
        ▼
WindowState
```

O `WindowContext` interpreta o evento e atualiza o estado correspondente:

```go
func (c *WindowContext) HandleEvent(event Event) {
    switch e := event.(type) {
    case ToplevelConfigureEvent:
        c.State.SetSize(
            e.Width,
            e.Height,
        )
    }
}
```

Observe que o `WindowState` permanece completamente desacoplado do protocolo. Ele apenas recebe uma alteração de estado, sem qualquer conhecimento sobre a origem dessa atualização.

---

# Conclusão

O `WindowContext` funciona como uma camada de agregação e tradução semântica entre o protocolo Wayland e o modelo interno da biblioteca.

Enquanto os proxies permanecem responsáveis apenas pela comunicação de baixo nível, o contexto reúne diversos objetos do protocolo em uma única entidade lógica capaz de representar uma janela completa.

O `Dispatcher`, por sua vez, limita-se a associar cada `object_id` ao `WindowContext` correspondente, permitindo que múltiplas janelas coexistam simultaneamente sem recorrer a estado global.

Como resultado, cada componente possui uma responsabilidade bem definida:

| Componente          | Responsabilidade                                                      |
| ------------------- | --------------------------------------------------------------------- |
| **Proxies Wayland** | Comunicação com o protocolo e tradução entre mensagens e eventos.     |
| **Dispatcher**      | Roteamento de eventos para o contexto correto.                        |
| **WindowContext**   | Composição dos objetos Wayland e interpretação semântica dos eventos. |
| **WindowState**     | Representação agnóstica do estado da janela.                          |

Essa separação reduz o acoplamento entre protocolo e modelo interno, melhora a reutilização dos proxies e facilita a implementação de novos backends sem alterar a representação da janela dentro da biblioteca.