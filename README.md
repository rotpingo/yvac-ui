# README

## About

This is the official Wails Vanilla-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.


## Windows Runtime Requirement

This app uses the [WebView2 Fixed Version Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/) for Windows builds.

To build:
1. Download the correct fixed runtime `.cab` from Microsoft.
2. Extract `msedgewebview2.exe` into a `webview2/` folder in the project root.
3. Ensure `wails.json` points to it:

```json
{
  "windows": {
    "webview2": {
      "runtime": "fixed",
      "path": "./webview2"
    }
  }
}
