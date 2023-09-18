use extism::{Manifest, Plugin};
use extism_manifest::Wasm;

fn main() {
    let wasm = Wasm::file("../plugin/http.wasm".to_owned());
    let manifest = Manifest::new([wasm])
        .with_allowed_host("jsonplaceholder.*.com");

    let mut plugin = Plugin::new_with_manifest(&manifest, [], true).unwrap();

    let data : String = plugin.call("_start", "").unwrap();
    println!("output: {}", data)
}
