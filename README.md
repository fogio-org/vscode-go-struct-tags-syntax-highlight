# Inline Go Struct Tags Syntax Highlight Extension for VS Code

Highlight golang struct tags according to your theme

[VS Code Marketplace](https://marketplace.visualstudio.com/items?itemName=fogio.inline-go-struct-tags-syntax-highlight):

![VS Code Marketplace Version](https://vsmarketplacebadges.dev/version-short/fogio.inline-go-struct-tags-syntax-highlight.svg?style=for-the-badge&colorA=555555&colorB=007ec6&label=VERSION)&nbsp;
![VS Code Marketplace Rating](https://vsmarketplacebadges.dev/rating-short/fogio.inline-go-struct-tags-syntax-highlight.svg?style=for-the-badge&colorA=555555&colorB=007ec6&label=RATING)&nbsp;
![VS Code Marketplace Downloads](https://vsmarketplacebadges.dev/downloads-short/fogio.inline-go-struct-tags-syntax-highlight.svg?style=for-the-badge&colorA=555555&colorB=007ec6&label=DOWNLOADS)&nbsp;
![VS Code Marketplace Installs](https://vsmarketplacebadges.dev/installs-short/fogio.inline-go-struct-tags-syntax-highlight.svg?style=for-the-badge&colorA=555555&colorB=007ec6&label=INSTALLS)

[Open VSX](https://open-vsx.org/extension/fogio/inline-go-struct-tags-syntax-highlight):

![Open VSX Version](https://img.shields.io/open-vsx/v/fogio/inline-go-struct-tags-syntax-highlight?style=for-the-badge&color=%23c260ef&label=VERSION)&nbsp;
![Open VSX Rating](https://img.shields.io/open-vsx/rating/fogio/inline-go-struct-tags-syntax-highlight?style=for-the-badge&color=%23c260ef&label=RATING)&nbsp;
![Open VSX Downloads](https://img.shields.io/open-vsx/dt/fogio/inline-go-struct-tags-syntax-highlight?style=for-the-badge&color=%23c260ef&label=DOWNLOADS)&nbsp;
![Open VSX Release Date](https://img.shields.io/open-vsx/release-date/fogio/inline-go-struct-tags-syntax-highlight?style=for-the-badge&color=%23c260ef&label=RELEASE%20DATE)


## Preview

Your theme's default colors

![preview-1](https://raw.githubusercontent.com/fogio-org/vscode-go-struct-tags-syntax-highlight/master/assets/img/preview-1.png)

![preview-2](https://raw.githubusercontent.com/fogio-org/vscode-go-struct-tags-syntax-highlight/master/assets/img/preview-2.png)

Customization

![preview-3](https://raw.githubusercontent.com/fogio-org/vscode-go-struct-tags-syntax-highlight/master/assets/img/preview-3.png)

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

![custom-font-style](https://raw.githubusercontent.com/fogio-org/vscode-go-struct-tags-syntax-highlight/master/assets/img/custom-font-style.png)

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
        "fontStyle": "italic"
      }
    }
  ]
}
```

### Font color

![custom-font-color](https://raw.githubusercontent.com/fogio-org/vscode-go-struct-tags-syntax-highlight/master/assets/img/custom-font-color.png)

```json
"editor.tokenColorCustomizations": {
  "textMateRules": [
    {
      "scope": "meta.struct-tag.pair.go support.type.property-name.json", // struct tag key
      "settings": {
        "foreground": "#9876AA",
      }
    },
    {
      "scope": "meta.struct-tag.pair.go punctuation.separator.dictionary.key-value.json", // struct tag separator (colon)
      "settings": {
        "foreground": "#BCBEC4"
      }
    },
    {
      "scope": "meta.struct-tag.pair.go string.quoted.double.json", // struct tag value
      "settings": {
        "foreground": "#6A8759",
      }
    }
  ]
}
```

### Font style with color

![custom-font-style-with-color](https://raw.githubusercontent.com/fogio-org/vscode-go-struct-tags-syntax-highlight/master/assets/img/custom-font-style-with-color.png)

```json
"editor.tokenColorCustomizations": {
  "textMateRules": [
    {
      "scope": "meta.struct-tag.pair.go support.type.property-name.json", // struct tag key
      "settings": {
        "foreground": "#9876AA",
        "fontStyle": "bold"
      }
    },
    {
      "scope": "meta.struct-tag.pair.go punctuation.separator.dictionary.key-value.json", // struct tag separator (colon)
      "settings": {
        "foreground": "#BCBEC4"
      }
    },
    {
      "scope": "meta.struct-tag.pair.go string.quoted.double.json", // struct tag value
      "settings": {
        "foreground": "#6A8759",
        "fontStyle": "italic"
      }
    }
  ]
}
```
