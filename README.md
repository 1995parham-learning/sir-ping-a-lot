# sir-ping-a-lot 🎺

<p align="center">
  <a href="https://github.com/1995parham-learning/sir-ping-a-lot/actions/workflows/ci.yml"><img alt="ci" src="https://github.com/1995parham-learning/sir-ping-a-lot/actions/workflows/ci.yml/badge.svg"></a>
  <img alt="go version" src="https://img.shields.io/badge/go-1.26-00ADD8?logo=go&logoColor=white">
  <img alt="license" src="https://img.shields.io/badge/license-GPL--3.0-blue">
  <img alt="last commit" src="https://img.shields.io/github/last-commit/1995parham-learning/sir-ping-a-lot">
  <img alt="repo size" src="https://img.shields.io/github/repo-size/1995parham-learning/sir-ping-a-lot">
</p>

> *"I like big checks and I cannot lie."*

A microservice-based **HTTP monitoring system**. You hand it URLs, and a small fleet of Go services keeps knighting around asking your endpoints the only question that matters: **are you still alive?**

## Architecture

<p align="center">
  <img alt="architecture" src="architecture/microservice.png" />
</p>

The system is split into small services that talk over [NATS](https://nats.io):

| Directory        | Role |
| ---------------- | ---- |
| [`user/`](user/)         | Public API — `Register`, `Login`, and adding URLs to monitor. |
| [`server/`](server/)     | Reads the URL table periodically and publishes each URL that needs checking to NATS. |
| [`checker/`](checker/)   | Subscribes to NATS, checks each URL's status, and publishes the result back. Runs as many instances. |
| [`saver/`](saver/)       | Bootstraps the database tables and persists the status results. |
| [`docs/`](docs/)         | The original project definition (LaTeX). |
| [`architecture/`](architecture/) | Architecture overview, diagram, and a top-level `docker-compose`. |


## Credits

Originally built by [Elaheh Dastan](https://github.com/elaheh-dastan), with mentoring from [Parham Alvani](https://github.com/1995parham). Consolidated here for learning purposes.

## License

Licensed under the [GNU General Public License v3.0](LICENSE).
