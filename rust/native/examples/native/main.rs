use virtblocks_rust_native::devices;
use virtblocks_rust_native::util;

fn main() {
    let file_name = util::build_file_name("guest", ".xml");
    println!("{}", file_name);

    let mut memballoon = devices::Memballoon::new();

    println!("{}", memballoon);

    memballoon.set_model(devices::MemballoonModel::Virtio);
    println!("{}", memballoon);
}
