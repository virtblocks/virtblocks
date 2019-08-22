extern crate virtblocks;

use virtblocks::devices::Memballoon;
use virtblocks::devices::MemballoonModel;

fn main() {
    let file_name = virtblocks::util::build_file_name("guest", ".xml");
    println!("{}", file_name);

    let mut memballoon = Memballoon::new();

    println!("{}", memballoon);

    memballoon.set_model(MemballoonModel::Virtio);
    println!("{}", memballoon);
}
