// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/type.proto

package org.mojolang.mojo.lang;

public interface AttributeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.Attribute)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  boolean hasStartPosition();
  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  org.mojolang.mojo.lang.Position getStartPosition();
  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getStartPositionOrBuilder();

  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  boolean hasEndPosition();
  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  org.mojolang.mojo.lang.Position getEndPosition();
  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getEndPositionOrBuilder();

  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>string package = 6;</code>
   */
  java.lang.String getPackage();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>string package = 6;</code>
   */
  com.google.protobuf.ByteString
      getPackageBytes();

  /**
   * <pre>
   *&#47; the package which this attribute belong to
   * </pre>
   *
   * <code>string name = 10;</code>
   */
  java.lang.String getName();
  /**
   * <pre>
   *&#47; the package which this attribute belong to
   * </pre>
   *
   * <code>string name = 10;</code>
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>repeated .mojo.lang.GenericArgument generic_arguments = 14;</code>
   */
  java.util.List<org.mojolang.mojo.lang.GenericArgument> 
      getGenericArgumentsList();
  /**
   * <code>repeated .mojo.lang.GenericArgument generic_arguments = 14;</code>
   */
  org.mojolang.mojo.lang.GenericArgument getGenericArguments(int index);
  /**
   * <code>repeated .mojo.lang.GenericArgument generic_arguments = 14;</code>
   */
  int getGenericArgumentsCount();
  /**
   * <code>repeated .mojo.lang.GenericArgument generic_arguments = 14;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.GenericArgumentOrBuilder> 
      getGenericArgumentsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.GenericArgument generic_arguments = 14;</code>
   */
  org.mojolang.mojo.lang.GenericArgumentOrBuilder getGenericArgumentsOrBuilder(
      int index);

  /**
   * <code>repeated .mojo.lang.Expression expressions = 15;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Expression> 
      getExpressionsList();
  /**
   * <code>repeated .mojo.lang.Expression expressions = 15;</code>
   */
  org.mojolang.mojo.lang.Expression getExpressions(int index);
  /**
   * <code>repeated .mojo.lang.Expression expressions = 15;</code>
   */
  int getExpressionsCount();
  /**
   * <code>repeated .mojo.lang.Expression expressions = 15;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.ExpressionOrBuilder> 
      getExpressionsOrBuilderList();
  /**
   * <code>repeated .mojo.lang.Expression expressions = 15;</code>
   */
  org.mojolang.mojo.lang.ExpressionOrBuilder getExpressionsOrBuilder(
      int index);
}
