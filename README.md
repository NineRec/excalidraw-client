# Excalidraw-client

Desktop client for [excalidraw](https://github.com/excalidraw/excalidraw) using ~~[tauri](https://github.com/tauri-apps/tauri)~~ [wails](https://github.com/wailsapp/wails), with custom font(中文字体).

## Local build/install

* install [wails](https://wails.app/gettingstarted/);
* run `wails dev` in the project directory while developing(a dev server that runs on http://localhost:34115);
* to build a redistributable, production mode package, use `wails build`.

## Custom font

In order to support handwritten chinese font, the font "Virgil.woff2" under `frontend/public/excalidraw-assets` & `frontend/public/excalidraw-assets-dev` is merged from origin "Virgil.woff2" and "XiaolaiSC-Regular.ttf" using fonttools. 

```
fonttools merge Virgil.ttf XiaolaiSC-Regular.ttf
```

And the origin font is renamed as "Virgil.woff2.bak" in same folder.

## TODO

- [ ] support import/export;

## Reference

* [Why Golang instead of Rust to develop the Krater desktop app](https://blog.moonguard.dev/why-golang-instead-of-rust-to-develop-the-krater-desktop-app)
* [nepaul - 300secs-create-a-desktop-application-tauri](https://github.com/nepaul/learn-to-code-by-coding/tree/b95c64ff30540278991c8bc613f8b2ca5dc9b31a/300secs-create-a-desktop-application-tauri)
