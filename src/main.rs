extern crate virtblocks;

fn main() {
    let file_name = virtblocks::util::build_file_name("guest", ".xml");
    println!("{}", file_name);
}
