// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/package_decl.proto

package org.mojolang.mojo.lang;

public final class PackageDeclProto {
  private PackageDeclProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_lang_PackageDecl_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_lang_PackageDecl_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\034mojo/lang/package_decl.proto\022\tmojo.lan" +
      "g\032\032mojo/lang/expression.proto\032\030mojo/lang" +
      "/position.proto\032\030mojo/lang/document.prot" +
      "o\"\311\001\n\013PackageDecl\022+\n\016start_position\030\001 \001(" +
      "\0132\023.mojo.lang.Position\022)\n\014end_position\030\002" +
      " \001(\0132\023.mojo.lang.Position\022%\n\010document\030\004 " +
      "\001(\0132\023.mojo.lang.Document\022\014\n\004name\030\n \001(\t\022-" +
      "\n\007package\030\013 \001(\0132\034.mojo.lang.ObjectLitera" +
      "lExprB]\n\026org.mojolang.mojo.langB\020Package" +
      "DeclProtoP\001Z/github.com/mojo-lang/lang/g" +
      "olang/mojo/lang;langb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.lang.ExpressionProto.getDescriptor(),
          org.mojolang.mojo.lang.PositionProto.getDescriptor(),
          org.mojolang.mojo.lang.DocumentProto.getDescriptor(),
        }, assigner);
    internal_static_mojo_lang_PackageDecl_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_lang_PackageDecl_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_lang_PackageDecl_descriptor,
        new java.lang.String[] { "StartPosition", "EndPosition", "Document", "Name", "Package", });
    org.mojolang.mojo.lang.ExpressionProto.getDescriptor();
    org.mojolang.mojo.lang.PositionProto.getDescriptor();
    org.mojolang.mojo.lang.DocumentProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
