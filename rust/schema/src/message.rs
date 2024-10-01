use ixc_message_api::header::MessageSelector;
use crate::codec::Codec;
use crate::SchemaValue;
use crate::structs::StructSchema;
use crate::value::ResponseValue;

#[non_exhaustive]
#[derive(Debug, Clone, Eq, PartialEq)]
pub struct MessageDescriptor<'a> {
    pub request_type: &'a str,
    pub response_type: &'a str,
    pub error_type: &'a str,
    pub events: &'a [&'a str],
}

pub trait Message<'a>: SchemaValue<'a> + StructSchema {
    const SELECTOR: MessageSelector;
    type Response<'b>: ResponseValue<'b>;
    type Error: ResponseValue<'static>; // TODO just accept error codes
    type Codec: Codec;
}
