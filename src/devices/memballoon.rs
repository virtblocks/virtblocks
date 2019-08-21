use std::fmt;

/// Type of balloon device
#[repr(C)]
#[derive(Copy, Clone)]
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

/// A balloon device
///
/// # Examples
///
/// ```
/// use virtblocks::devices::Memballoon;
/// use virtblocks::devices::MemballoonModel;
///
/// let mut memballoon = Memballoon::new();
///
/// assert_eq!("", memballoon.to_string());
///
/// memballoon.set_model(MemballoonModel::Virtio);
/// assert_eq!("virtio-memballoon", memballoon.to_string());
/// ```
#[derive(Default)]
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

    pub fn get_model(&self) -> MemballoonModel {
        self.model
    }
}

impl fmt::Display for Memballoon {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let model_str = match self.model {
            MemballoonModel::Virtio => "virtio-memballoon",
            MemballoonModel::VirtioNonTransitional => "virtio-memballoon-non-transitional",
            MemballoonModel::VirtioTransitional => "virtio-memballoon-transitional",
            MemballoonModel::None => "",
        };
        write!(f, "{}", model_str)
    }
}
