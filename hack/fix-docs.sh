DOCS_PATH=".work/rancher/rancher2/docs/resources"

# inject description if missing in the frontmatter
for file in "${DOCS_PATH}"/*.md; do
    if ! sed -n '/^---/,/^---/p' "$file" | grep -q "description:"; then
        echo "Injecting missing description into $file"
        # Inject the description line after the opening frontmatter ---
        sed -i '2i description: "Managed resource"' "$file"
    fi
done

# fix broken examples

# [<FOO>] to ["<FOO>"]
find "${DOCS_PATH}" -name "*.md" -type f -exec sed -i -E 's/(\[)(<[^>]+>)(\])/\1"\2"\3/g' {} +

# <FOO> to "<FOO>"
find "${DOCS_PATH}" -name "*.md" -type f -exec sed -i -E 's/(= )(<[^>]+>)/\1"\2"/g' {} +
