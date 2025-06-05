# vscode-go-struct-tags-syntax-highlight

Inline Go Struct Tags Syntax Highlight Extension for VS Code

## Disabling Semantic Highlighting (for use with standard themes)

By default, standard VS Code themes use semantic highlighting, which overrides custom struct tag highlighting. To use this extension with standard themes, you need to disable semantic highlighting for Go files:

### Option 1: Disable only for Go

```json
"[go]": {
  "editor.semanticHighlighting.enabled": false
}
```

### Option 2: Disable globally

```json
"editor.semanticHighlighting.enabled": false
```

Add this to your `settings.json` (open Command Palette → Preferences: Open Settings (JSON)).

---

## Customizing Colors in User Settings

You can customize the colors for struct tag keys and values using `editor.tokenColorCustomizations` in your `settings.json`:

```json
"editor.tokenColorCustomizations": {
  "textMateRules": [
    {
      "scope": "entity.name.tag.go",
      "settings": {
        "foreground": "#FF8800"
      }
    },
    {
      "scope": "string.quoted.double.go",
      "settings": {
        "foreground": "#00AAFF"
      }
    }
  ]
}
```
