# Go Echo Example with Pongo2 Template Engine

## Template

It seems Golang template from the standard library is totally unusable for backend. It works only when you want to render very simple HTML text. On server-side, however, your templates are usually very complicated, with messy inheritance and reuse.

When I try to implement the inheritance following this article [Golang Template Inheritance](https://siongui.github.io/2017/02/05/go-template-inheritance-jinja2-extends-include/), its problem arises. When you `define` the same `block` in multiple template files, and precompile them all at startup, only the last defined block wins. So you cannot implement the template inheritance like Jinja2.

I guess the problem is all files and blocks are parsed into the same namespace so you cannot `define` the same `block` twice. The template engine is not usable on server-side.

See [this article](https://machiel.me/post/pongo2-with-echo-or-net-http/) explains the Golang template package cannot meet backend needs.