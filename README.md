# Budgie - WIP

This is very much so a Work in Progress.

## Background

I wanted to take an existing process and port that into application.

So that I stole a Google Sheets template on Budgeting and I'm turning that into an app.

## Goals

My main goal here is really to explore and get better at golang. The secondary goal is to improve on "Professional" Software Developer skills, such as testing and logging.

That's why I picked golang for the backend and went with more familiar technologies for the frontend.

## Built With

### Backend

- go
- chi-router
- Clerk Go SDK
- PostgreSQL via Docker
- go-migrate

Note: gofiber was originally here, but it's incompatible with Clerk as it does not use the Standard Library `net/http` package.

### Frontend

- TypeScript
- Next.js
- React
- TailwindCSS
- Clerk Auth - Big Believer of Avoiding "Reinventing the Wheel"
- [TODO] Vitest
- [TODO] React Testing Library
- [TODO] Mock Service Worker

### Other

- Makefile to automate development enviroment setup.
