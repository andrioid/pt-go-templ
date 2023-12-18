# Prototype Stack

## Motivation

As frontend grows in complexity, we're reminded of simpler times with server-rendered HTML. At the same time I've grown to really love the composability of React components.

My hope is that this is a viable web stack, where I can put all of the domain logic in one place. Regardless of it's a first render, or responding to a request.

## Stack

- Go
- Tailwind
- [Templ](https://templ.guide)
- HTMX
- Repository pattern
- SQLite

## Tech evaluations

- [x] Evaluate routers
  - [x] ~echo~: Seems to work fine, nothing special though
  - [x] ServeMux: Getting better in Go 1.22 - trying to keep this minimal
- [x] Evaluate various database layers
  - [x] ~Ent~: If was building an enterprise app, yes
  - [x] ~Gorm~: Good/Bad: 80/20. The bad is too high.
  - [x] Manual: Works, minimal. We're going with minimal.
- [ ] When HTMX isn't enough?
  - [ ] Hyperscript?

## Tasks

- Play more with templ
- Add SQLite
- Read items from database
- Update items in database
- Serve /public for CSS
- Embed assets into binary
- Experiment with SQLite, maybe Gorm

## Global install deps

- go install github.com/pressly/goose/v3/cmd/goose@latest
- go install github.com/a-h/templ/cmd/templ@latest
