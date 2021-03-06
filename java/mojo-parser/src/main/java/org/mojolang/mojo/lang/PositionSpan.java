// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/position.proto

package org.mojolang.mojo.lang;

/**
 * Protobuf type {@code mojo.lang.PositionSpan}
 */
public  final class PositionSpan extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.lang.PositionSpan)
    PositionSpanOrBuilder {
private static final long serialVersionUID = 0L;
  // Use PositionSpan.newBuilder() to construct.
  private PositionSpan(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private PositionSpan() {
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private PositionSpan(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            org.mojolang.mojo.lang.Position.Builder subBuilder = null;
            if (start_ != null) {
              subBuilder = start_.toBuilder();
            }
            start_ = input.readMessage(org.mojolang.mojo.lang.Position.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(start_);
              start_ = subBuilder.buildPartial();
            }

            break;
          }
          case 18: {
            org.mojolang.mojo.lang.Position.Builder subBuilder = null;
            if (end_ != null) {
              subBuilder = end_.toBuilder();
            }
            end_ = input.readMessage(org.mojolang.mojo.lang.Position.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(end_);
              end_ = subBuilder.buildPartial();
            }

            break;
          }
          default: {
            if (!parseUnknownFieldProto3(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.lang.PositionProto.internal_static_mojo_lang_PositionSpan_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.lang.PositionProto.internal_static_mojo_lang_PositionSpan_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.lang.PositionSpan.class, org.mojolang.mojo.lang.PositionSpan.Builder.class);
  }

  public static final int START_FIELD_NUMBER = 1;
  private org.mojolang.mojo.lang.Position start_;
  /**
   * <code>.mojo.lang.Position start = 1;</code>
   */
  public boolean hasStart() {
    return start_ != null;
  }
  /**
   * <code>.mojo.lang.Position start = 1;</code>
   */
  public org.mojolang.mojo.lang.Position getStart() {
    return start_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : start_;
  }
  /**
   * <code>.mojo.lang.Position start = 1;</code>
   */
  public org.mojolang.mojo.lang.PositionOrBuilder getStartOrBuilder() {
    return getStart();
  }

  public static final int END_FIELD_NUMBER = 2;
  private org.mojolang.mojo.lang.Position end_;
  /**
   * <code>.mojo.lang.Position end = 2;</code>
   */
  public boolean hasEnd() {
    return end_ != null;
  }
  /**
   * <code>.mojo.lang.Position end = 2;</code>
   */
  public org.mojolang.mojo.lang.Position getEnd() {
    return end_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : end_;
  }
  /**
   * <code>.mojo.lang.Position end = 2;</code>
   */
  public org.mojolang.mojo.lang.PositionOrBuilder getEndOrBuilder() {
    return getEnd();
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (start_ != null) {
      output.writeMessage(1, getStart());
    }
    if (end_ != null) {
      output.writeMessage(2, getEnd());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (start_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getStart());
    }
    if (end_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getEnd());
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojolang.mojo.lang.PositionSpan)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.lang.PositionSpan other = (org.mojolang.mojo.lang.PositionSpan) obj;

    boolean result = true;
    result = result && (hasStart() == other.hasStart());
    if (hasStart()) {
      result = result && getStart()
          .equals(other.getStart());
    }
    result = result && (hasEnd() == other.hasEnd());
    if (hasEnd()) {
      result = result && getEnd()
          .equals(other.getEnd());
    }
    result = result && unknownFields.equals(other.unknownFields);
    return result;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (hasStart()) {
      hash = (37 * hash) + START_FIELD_NUMBER;
      hash = (53 * hash) + getStart().hashCode();
    }
    if (hasEnd()) {
      hash = (37 * hash) + END_FIELD_NUMBER;
      hash = (53 * hash) + getEnd().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.lang.PositionSpan parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.lang.PositionSpan parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.mojolang.mojo.lang.PositionSpan prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code mojo.lang.PositionSpan}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.lang.PositionSpan)
      org.mojolang.mojo.lang.PositionSpanOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.lang.PositionProto.internal_static_mojo_lang_PositionSpan_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.lang.PositionProto.internal_static_mojo_lang_PositionSpan_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.lang.PositionSpan.class, org.mojolang.mojo.lang.PositionSpan.Builder.class);
    }

    // Construct using org.mojolang.mojo.lang.PositionSpan.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      if (startBuilder_ == null) {
        start_ = null;
      } else {
        start_ = null;
        startBuilder_ = null;
      }
      if (endBuilder_ == null) {
        end_ = null;
      } else {
        end_ = null;
        endBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.lang.PositionProto.internal_static_mojo_lang_PositionSpan_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.PositionSpan getDefaultInstanceForType() {
      return org.mojolang.mojo.lang.PositionSpan.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.PositionSpan build() {
      org.mojolang.mojo.lang.PositionSpan result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.lang.PositionSpan buildPartial() {
      org.mojolang.mojo.lang.PositionSpan result = new org.mojolang.mojo.lang.PositionSpan(this);
      if (startBuilder_ == null) {
        result.start_ = start_;
      } else {
        result.start_ = startBuilder_.build();
      }
      if (endBuilder_ == null) {
        result.end_ = end_;
      } else {
        result.end_ = endBuilder_.build();
      }
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return (Builder) super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return (Builder) super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return (Builder) super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return (Builder) super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return (Builder) super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return (Builder) super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.mojolang.mojo.lang.PositionSpan) {
        return mergeFrom((org.mojolang.mojo.lang.PositionSpan)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.lang.PositionSpan other) {
      if (other == org.mojolang.mojo.lang.PositionSpan.getDefaultInstance()) return this;
      if (other.hasStart()) {
        mergeStart(other.getStart());
      }
      if (other.hasEnd()) {
        mergeEnd(other.getEnd());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      org.mojolang.mojo.lang.PositionSpan parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.lang.PositionSpan) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private org.mojolang.mojo.lang.Position start_ = null;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> startBuilder_;
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    public boolean hasStart() {
      return startBuilder_ != null || start_ != null;
    }
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    public org.mojolang.mojo.lang.Position getStart() {
      if (startBuilder_ == null) {
        return start_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : start_;
      } else {
        return startBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    public Builder setStart(org.mojolang.mojo.lang.Position value) {
      if (startBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        start_ = value;
        onChanged();
      } else {
        startBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    public Builder setStart(
        org.mojolang.mojo.lang.Position.Builder builderForValue) {
      if (startBuilder_ == null) {
        start_ = builderForValue.build();
        onChanged();
      } else {
        startBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    public Builder mergeStart(org.mojolang.mojo.lang.Position value) {
      if (startBuilder_ == null) {
        if (start_ != null) {
          start_ =
            org.mojolang.mojo.lang.Position.newBuilder(start_).mergeFrom(value).buildPartial();
        } else {
          start_ = value;
        }
        onChanged();
      } else {
        startBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    public Builder clearStart() {
      if (startBuilder_ == null) {
        start_ = null;
        onChanged();
      } else {
        start_ = null;
        startBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    public org.mojolang.mojo.lang.Position.Builder getStartBuilder() {
      
      onChanged();
      return getStartFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    public org.mojolang.mojo.lang.PositionOrBuilder getStartOrBuilder() {
      if (startBuilder_ != null) {
        return startBuilder_.getMessageOrBuilder();
      } else {
        return start_ == null ?
            org.mojolang.mojo.lang.Position.getDefaultInstance() : start_;
      }
    }
    /**
     * <code>.mojo.lang.Position start = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> 
        getStartFieldBuilder() {
      if (startBuilder_ == null) {
        startBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder>(
                getStart(),
                getParentForChildren(),
                isClean());
        start_ = null;
      }
      return startBuilder_;
    }

    private org.mojolang.mojo.lang.Position end_ = null;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> endBuilder_;
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    public boolean hasEnd() {
      return endBuilder_ != null || end_ != null;
    }
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    public org.mojolang.mojo.lang.Position getEnd() {
      if (endBuilder_ == null) {
        return end_ == null ? org.mojolang.mojo.lang.Position.getDefaultInstance() : end_;
      } else {
        return endBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    public Builder setEnd(org.mojolang.mojo.lang.Position value) {
      if (endBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        end_ = value;
        onChanged();
      } else {
        endBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    public Builder setEnd(
        org.mojolang.mojo.lang.Position.Builder builderForValue) {
      if (endBuilder_ == null) {
        end_ = builderForValue.build();
        onChanged();
      } else {
        endBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    public Builder mergeEnd(org.mojolang.mojo.lang.Position value) {
      if (endBuilder_ == null) {
        if (end_ != null) {
          end_ =
            org.mojolang.mojo.lang.Position.newBuilder(end_).mergeFrom(value).buildPartial();
        } else {
          end_ = value;
        }
        onChanged();
      } else {
        endBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    public Builder clearEnd() {
      if (endBuilder_ == null) {
        end_ = null;
        onChanged();
      } else {
        end_ = null;
        endBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    public org.mojolang.mojo.lang.Position.Builder getEndBuilder() {
      
      onChanged();
      return getEndFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    public org.mojolang.mojo.lang.PositionOrBuilder getEndOrBuilder() {
      if (endBuilder_ != null) {
        return endBuilder_.getMessageOrBuilder();
      } else {
        return end_ == null ?
            org.mojolang.mojo.lang.Position.getDefaultInstance() : end_;
      }
    }
    /**
     * <code>.mojo.lang.Position end = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder> 
        getEndFieldBuilder() {
      if (endBuilder_ == null) {
        endBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.lang.Position, org.mojolang.mojo.lang.Position.Builder, org.mojolang.mojo.lang.PositionOrBuilder>(
                getEnd(),
                getParentForChildren(),
                isClean());
        end_ = null;
      }
      return endBuilder_;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFieldsProto3(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:mojo.lang.PositionSpan)
  }

  // @@protoc_insertion_point(class_scope:mojo.lang.PositionSpan)
  private static final org.mojolang.mojo.lang.PositionSpan DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.lang.PositionSpan();
  }

  public static org.mojolang.mojo.lang.PositionSpan getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<PositionSpan>
      PARSER = new com.google.protobuf.AbstractParser<PositionSpan>() {
    @java.lang.Override
    public PositionSpan parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new PositionSpan(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<PositionSpan> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<PositionSpan> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.lang.PositionSpan getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

