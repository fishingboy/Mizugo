// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: msgid.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
/// <summary>Holder for reflection information generated from msgid.proto</summary>
public static partial class MsgidReflection {

  #region Descriptor
  /// <summary>File descriptor for msgid.proto</summary>
  public static pbr::FileDescriptor Descriptor {
    get { return descriptor; }
  }
  private static pbr::FileDescriptor descriptor;

  static MsgidReflection() {
    byte[] descriptorData = global::System.Convert.FromBase64String(
        string.Concat(
          "Cgttc2dpZC5wcm90bypuCgVNc2dJRBILCgdVbmtub3duEAASCAoES2V5URAB",
          "EggKBEtleUEQAhIJCgVKc29uURADEgkKBUpzb25BEAQSCgoGUHJvdG9REAUS",
          "CgoGUHJvdG9BEAYSCgoGUExpc3RREAcSCgoGUExpc3RBEAhCDFoKL21zZ3M7",
          "bXNnc2IGcHJvdG8z"));
    descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
        new pbr::FileDescriptor[] { },
        new pbr::GeneratedClrTypeInfo(new[] {typeof(global::MsgID), }, null, null));
  }
  #endregion

}
#region Enums
/// <summary>
/// 訊息編號
/// </summary>
public enum MsgID {
  /// <summary>
  /// 不明/錯誤訊息編號, 此編號不可使用
  /// </summary>
  [pbr::OriginalName("Unknown")] Unknown = 0,
  /// <summary>
  /// 要求密鑰
  /// </summary>
  [pbr::OriginalName("KeyQ")] KeyQ = 1,
  /// <summary>
  /// 回應密鑰
  /// </summary>
  [pbr::OriginalName("KeyA")] KeyA = 2,
  /// <summary>
  /// 要求Json
  /// </summary>
  [pbr::OriginalName("JsonQ")] JsonQ = 3,
  /// <summary>
  /// 回應Json
  /// </summary>
  [pbr::OriginalName("JsonA")] JsonA = 4,
  /// <summary>
  /// 要求Proto
  /// </summary>
  [pbr::OriginalName("ProtoQ")] ProtoQ = 5,
  /// <summary>
  /// 回應Proto
  /// </summary>
  [pbr::OriginalName("ProtoA")] ProtoA = 6,
  /// <summary>
  /// 要求PList
  /// </summary>
  [pbr::OriginalName("PListQ")] PlistQ = 7,
  /// <summary>
  /// 回應PList
  /// </summary>
  [pbr::OriginalName("PListA")] PlistA = 8,
}

#endregion


#endregion Designer generated code
