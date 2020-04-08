# markymark

Simple tool to make nested .md files possible.

It replaces placeholders in the form of `{{ path/to/file.md }}` with the actual content of this file.

Example:

```
This is my markdown file with a couple of placeholders

## Section 1 
{{ ./snippets/myFirstFile.md }}


## Section 2
{{ ./snippets/myOtherFile.md }}

```

Usage:

```
./markymark myfile.md
```

Will produce a new file `myfile_gen.md` where all placeholders are replaced with the actual file contents of their file paths

