
# Static Site Generator

This project is a **Go-based static site generator** that leverages Go’s `html/template` package and the [Goldmark](https://github.com/yuin/goldmark) Markdown processor to produce a fast, maintainable, and extensible website.

## Key Features

1. **Go Templates**

   * Uses `html/template` for safe, easy-to-maintain layouts.
   * Supports base and page templates (`base.html`, `page.html`), partials, and custom functions.

2. **Markdown Content**

   * Parses `.md` files from the `content/` directory.
   * Converts Markdown (with GitHub Flavored Markdown extensions) into sanitized HTML pages.

3. **Database Integration**

   * Optional support for a lightweight database (SQLite / PostgreSQL).
   * Schema migration using [golang-migrate](https://github.com/golang-migrate/migrate).
   * Stores page metadata (title, slug, date, tags) and provides programmatic querying.

4. **Asset Pipeline**

   * Copies static assets (`static/` folder) into the `public/` output directory.
   * Serves CSS, JS, images, and other resources alongside generated HTML.

## Folder Structure

```text
static-site-generator/
├── content/             # Markdown pages
│   ├── index.md         # Home page
│   ├── about.md         # About page
│   └── site-description.md  # Project overview (this file)
├── templates/           # Go HTML templates
│   ├── base.html
│   └── page.html
├── static/              # Unprocessed assets
│   └── css/
│       └── style.css
├── db/                  # Database migrations and schema
│   ├── migrations/
│   └── seed.sql
├── public/              # Generated output
├── main.go              # Generator entry point
└── go.mod               # Module definition
```

## Usage

```bash
# Build the generator
go build -o bin/static-site-generator ./main.go

# Run with default settings
./bin/static-site-generator

# Custom content or template directories
./bin/static-site-generator -content=content -templates=templates -static=static -out=public
```

## Database Setup

1. **Configuration**

   * Define your database DSN in `config.yaml` or via environment variables.

2. **Migrations**

   ```bash
   migrate -path db/migrations -database "$DATABASE_URL" up
   ```

3. **Running**

   ```bash
   DATABASE_URL=postgres://user:pass@localhost:5432/site_db?sslmode=disable \
     bin/static-site-generator
   ```

This setup ensures your static content can be indexed, searched, or extended with dynamic features while preserving the speed and simplicity of a static site.
