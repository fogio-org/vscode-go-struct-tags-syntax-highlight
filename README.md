# Inline Go Struct Tags Syntax Highlight Extension for VS Code

Highlight golang struct tags according to your theme

## Preview

Your theme's default colors

...

Customization

...

## Install

By default, some VS Code themes use semantic highlighting, which overrides custom struct tag highlighting with a simple string. To use this extension with such themes, you need to disable semantic highlighting for Go files:

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

You can customize the colors for struct tag keys and values using `editor.tokenColorCustomizations` in your `settings.json`.

### Font style

...

```json
"editor.tokenColorCustomizations": {
  "textMateRules": [
    {
      "scope": "meta.struct-tag.pair.go support.type.property-name.json", // struct tag key
      "settings": {
        "fontStyle": "bold"
      }
    },
    {
      "scope": "meta.struct-tag.pair.go punctuation.separator.dictionary.key-value.json", // struct tag separator (colon)
      "settings": {
        "fontStyle": ""
      }
    },
    {
      "scope": "meta.struct-tag.pair.go string.quoted.double.json", // struct tag value
      "settings": {
        "fontStyle": "italic"
      }
    }
  ]
}
```

...

```json
"editor.tokenColorCustomizations": {
  "textMateRules": [
    {
      "scope": "meta.struct-tag.pair.go support.type.property-name.json", // struct tag key
      "settings": {
        "fontStyle": "bold underline"
      }
    },
    {
      "scope": "meta.struct-tag.pair.go punctuation.separator.dictionary.key-value.json", // struct tag separator (colon)
      "settings": {
        "fontStyle": ""
      }
    },
    {
      "scope": "meta.struct-tag.pair.go string.quoted.double.json", // struct tag value
      "settings": {
        "fontStyle": ""
      }
    }
  ]
}
```

### Font color

...

```json
"editor.tokenColorCustomizations": {
  "textMateRules": [
    {
      "scope": "meta.struct-tag.pair.go support.type.property-name.json", // struct tag key
      "settings": {
        "foreground": "#FF8800",
      }
    },
    {
      "scope": "meta.struct-tag.pair.go punctuation.separator.dictionary.key-value.json", // struct tag separator (colon)
      "settings": {
        "foreground": "#00AAFF"
      }
    },
    {
      "scope": "meta.struct-tag.pair.go string.quoted.double.json", // struct tag value
      "settings": {
        "foreground": "#FF00AA",
      }
    }
  ]
}
```

### Font style with color

...

```json
"editor.tokenColorCustomizations": {
  "textMateRules": [
    {
      "scope": "meta.struct-tag.pair.go support.type.property-name.json", // struct tag key
      "settings": {
        "foreground": "#FF8800",
        "fontStyle": "bold"
      }
    },
    {
      "scope": "meta.struct-tag.pair.go punctuation.separator.dictionary.key-value.json", // struct tag separator (colon)
      "settings": {
        "foreground": "#00AAFF"
      }
    },
    {
      "scope": "meta.struct-tag.pair.go string.quoted.double.json", // struct tag value
      "settings": {
        "foreground": "#FF00AA",
        "fontStyle": "italic"
      }
    }
  ]
}
```
