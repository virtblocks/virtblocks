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
pub struct Memballoon {
    model: MemballoonModel,
}

impl Memballoon {
    pub fn new() -> Self {
        Self {
            model: MemballoonModel::None,
        }
    }

    pub fn set_model(&mut self, model: MemballoonModel) {
        self.model = model;
    }

    pub fn get_model(&self) -> MemballoonModel {
        self.model
    }
}

impl Default for Memballoon {
    fn default() -> Self {
        Self::new()
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
