#![doc = include_str!(concat!(env!("CARGO_MANIFEST_DIR"), "/README.md"))]

#[doc(inline)]
pub use ixc_core::{Context, Result, EventBus, ensure, bail, fmt_error};
#[doc(inline)]
pub use ixc_core::resource::Resources;

#[doc(inline)]
pub use ixc_message_api::AccountID;
#[doc(inline)]
pub use ixc_schema::{SchemaValue, Str};
#[doc(inline)]
pub use state_objects::{Map, Item, Accumulator, AccumulatorMap};
#[doc(inline)]
pub use simple_time::{Time, Duration};

/// Re-exports for macros and other utilities.
pub mod internal {
    pub use ixc_core::*;
    pub use ixc_schema::binary::NativeBinaryCodec;
    pub use ixc_schema::codec::{decode_value};
}

#[cfg(feature = "core_macros")]
#[allow(unused_imports)]
#[macro_use]
extern crate ixc_core_macros;
#[cfg(feature = "core_macros")]
#[doc(inline)]
pub use ixc_core_macros::*;

#[cfg(feature = "schema_macros")]
#[allow(unused_imports)]
#[macro_use]
extern crate ixc_schema_macros;
#[cfg(feature = "schema_macros")]
#[doc(inline)]
pub use ixc_schema_macros::*;

#[cfg(feature = "state_objects_macros")]
#[allow(unused_imports)]
#[macro_use]
extern crate state_objects_macros;
#[cfg(feature = "state_objects_macros")]
#[doc(inline)]
pub use state_objects_macros::*;
