use std::fmt;

/// Type of balloon device
#[repr(C)]
#[derive(Copy, Clone, Debug)]
pub enum MemballoonModel {
    None,
    Virtio,
    VirtioNonTransitional,
    VirtioTransitional,
}

impl Default for MemballoonModel {
    fn default() -> Self {
        Self::None
    }
}

impl fmt::Display for MemballoonModel {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let model_str = match self {
            MemballoonModel::Virtio => "virtio-memballoon",
            MemballoonModel::VirtioNonTransitional => "virtio-memballoon-non-transitional",
            MemballoonModel::VirtioTransitional => "virtio-memballoon-transitional",
            MemballoonModel::None => "",
        };
        write!(f, "{}", model_str)
    }
}

/// A balloon device
///
/// # Examples
///
/// ```
/// use virtblocks_rust_native::devices;
///
/// let mut memballoon = devices::Memballoon::new();
///
/// assert_eq!("", memballoon.to_string());
///
/// memballoon.set_model(devices::MemballoonModel::Virtio);
/// assert_eq!("virtio-memballoon", memballoon.to_string());
/// ```
#[derive(Default, Debug)]
pub struct Memballoon {
    model: MemballoonModel,
}

impl Memballoon {
    pub fn new() -> Self {
        Default::default()
    }

    pub fn set_model(&mut self, model: MemballoonModel) {
        self.model = model;
    }

    pub fn model(&self) -> MemballoonModel {
        self.model
    }
}

impl fmt::Display for Memballoon {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.model)
    }
}
