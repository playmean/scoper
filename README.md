# scoper

🎯 Scoper is a tool for reports collection and manage.

---

[![CircleCI](https://circleci.com/gh/playmean/scoper.svg?style=shield)](https://circleci.com/gh/playmean/scoper) [![GPL-3.0 license](https://img.shields.io/github/license/playmean/scoper.svg)](https://github.com/playmean/scoper/blob/master/LICENSE)

## Getting Started

```bash
docker run -d -p 8080:8080 playmean/scoper:latest
```

Or if a custom configuration is needed:

```bash
docker run -d -p 8080:8080 -v /dir/with/cfg:/data playmean/scoper:latest
```

Will grab `config.json` from specified directory.
