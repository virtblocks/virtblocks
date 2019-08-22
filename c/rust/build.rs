extern crate libtool;

fn main() {
    libtool::generate_convenience_lib("libvirtblocks_c_rust").unwrap();
}
