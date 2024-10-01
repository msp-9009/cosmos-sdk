//! The raw handler and host backend interfaces.
use crate::code::ErrorCode;
use crate::packet::MessagePacket;

/// A handler for an account.
pub trait RawHandler {
    /// Handle a message packet.
    fn handle(&self, message_packet: &mut MessagePacket, callbacks: &dyn HostBackend) -> Result<(), HandlerErrorCode>;
}

pub use allocator_api2::alloc::Allocator;

/// A host backend for the handler.
pub trait HostBackend {
    /// Invoke a message packet.
    fn invoke(&self, message_packet: &mut MessagePacket, allocator: &dyn Allocator) -> Result<(), ErrorCode>;
    /// Allocate memory for a message response.
    /// The memory management expectation of handlers is that the caller
    /// deallocates both the memory it allocated and any memory allocated
    /// for the response by the callee.
    /// The alloc function in the host backend should return a pointer to
    /// memory that the caller knows how to free and such allocated
    /// memory should be referenced in the message packet's out pointers.
    fn allocator(&self) -> &dyn Allocator;
}

#[derive(Clone, Copy, PartialEq, Eq, Debug)]
/// An error code returned by a handler.
pub enum HandlerError {
    /// A known handler error code, usually returned by handler implementation libraries.
    KnownCode(HandlerErrorCode),
    /// A custom error code returned by a handler.
    Custom(u16),
}

/// A pre-defined error code that is usually returned by handler implementation libraries,
/// rather than handlers themselves.
#[derive(Clone, Copy, PartialEq, Eq, Debug)]
pub enum HandlerErrorCode {
    /// The handler doesn't handle the specified message.
    MessageNotHandled,
    /// Encoding error.
    EncodingError,
}
