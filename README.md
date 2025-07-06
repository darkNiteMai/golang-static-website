
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

---

## Updating Content

* **Markdown Pages:**  
  Add or edit `.md` files in the `content/` directory.  
  Each file represents a page (e.g., `about.md`, `index.md`).

* **Metadata:**  
  Include front matter (YAML or TOML) at the top of each Markdown file for title, date, tags, etc.

---

## Modifying Templates

* **Base Layout:**  
  Edit `templates/base.html` for site-wide layout (header, footer, etc.).

* **Page Templates:**  
  Update or add templates in `templates/` (e.g., `page.html`) to control how content is rendered.

* **Partials:**  
  Create reusable template snippets (e.g., `header.html`, `footer.html`) and include them in main templates.

---

## Managing Static Assets

* **CSS/JS/Images:**  
  Place static files in the `static/` directory.  
  Example: `static/css/style.css`

* **Referencing Assets:**  
  Use relative paths in your templates to reference static files.

---

## Database Integration

* **Migrations:**  
  Add migration files to `db/migrations/` for schema changes.

* **Seeding:**  
  Place initial data in `db/seed.sql`.

* **Configuration:**  
  Update database connection settings in `config.yaml`, environment variables, or `docker-compose.yml`.

* **Running Migrations:**  
  Use the `migrate` tool:

  ```sh
  migrate -path db/migrations -database "$DATABASE_URL" up
  ```

---

## Configuration

* **Environment Variables:**  
  Set variables such as `DATABASE_URL` for database connections.

* **Command-Line Flags:**  
  Override content, template, static, or output directories:

  ```sh
  ./bin/static-site-generator -content=content -templates=templates -static=static -out=public
  ```

---

## Docker Packaging

* **Dockerfile:**  
  Edit to change build steps or add dependencies.

* **docker-compose.yml:**  
  Update to configure services (app, database) and environment variables.

* **Build and Run:**

  ```sh
  docker-compose up --build
  ```

---

## Running Locally

* **Build the App:**

  ```sh
  go build -o bin/static-site-generator ./main.go
  ```

* **Run the App:**

  ```sh
  ./bin/static-site-generator
  ```

* **Access the Site:**  
  Open `http://localhost:8080` in your browser.

---

## Extending Functionality

* **Add Custom Template Functions:**  
  Register new functions in `main.go` for use in templates.

* **API Endpoints:**  
  Extend the Go application to serve dynamic data or APIs.

* **Search/Indexing:**  
  Integrate search features using the database or third-party services.

---

## Additional Resources

* [Go Templates Documentation](https://pkg.go.dev/html/template)
* [Goldmark Markdown Processor](https://github.com/yuin/goldmark)
* [Docker Documentation](https://docs.docker.com/)
* [golang-migrate](https://github.com/golang-migrate/migrate)
