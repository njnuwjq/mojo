// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/position.proto

package org.mojolang.mojo.lang;

public interface PositionOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.Position)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * filename, if any
   * </pre>
   *
   * <code>string filename = 1;</code>
   */
  java.lang.String getFilename();
  /**
   * <pre>
   * filename, if any
   * </pre>
   *
   * <code>string filename = 1;</code>
   */
  com.google.protobuf.ByteString
      getFilenameBytes();

  /**
   * <pre>
   * offset, starting at 0 (byte count)
   * </pre>
   *
   * <code>int32 offset = 2;</code>
   */
  int getOffset();

  /**
   * <pre>
   *&lt; line number, starting at 1
   * </pre>
   *
   * <code>int32 line = 3;</code>
   */
  int getLine();

  /**
   * <pre>
   *&lt; column number, starting at 1 (byte count)
   * </pre>
   *
   * <code>int32 column = 4;</code>
   */
  int getColumn();
}
