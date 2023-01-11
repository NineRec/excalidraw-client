# excalidraw-client

Desktop client for [excalidraw](https://github.com/excalidraw/excalidraw) using [tauri](https://github.com/tauri-apps/tauri), with custom font(中文字体).

## Local build/install

In order to build or run the code locally, you will need to have the ablility to run excalidraw(using yarn), and run tauri(using rust/cargo) locally.

For excalidraw, can refer [excalidraw#local-installation](https://github.com/excalidraw/excalidraw#local-installation).
For tauri, cna refer [TAURI - Integrate into Existing Project](https://tauri.app/v1/guides/getting-started/setup/integrate).

So, you need to have `Node.js, Yarn, Rust/Cargo, tauri-cli` installed, and then:

* run `cargo tauri build` can get you the release version of this app;
* run `cargo tauri dev` can get you debug version app running;

## known issues

- [ ] "Add to Excalidraw" from [Excalidraw Libraries](https://libraries.excalidraw.com/)
    - current solution: download first and then import to library;
- [ ] export image and save as PNG/SVG
    - current solution: export image to clipboard and then paste the image to other document;

## reference

* [nepaul- 300secs-create-a-desktop-application-tauri](https://github.com/nepaul/learn-to-code-by-coding/tree/b95c64ff30540278991c8bc613f8b2ca5dc9b31a/300secs-create-a-desktop-application-tauri)