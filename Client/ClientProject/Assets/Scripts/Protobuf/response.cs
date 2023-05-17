// <auto-generated>
//   This file was generated by a tool; you should avoid making direct changes.
//   Consider using 'partial classes' to extend these types
//   Input: response.proto
// </auto-generated>

#region Designer generated code
#pragma warning disable CS0612, CS0618, CS1591, CS3021, IDE0079, IDE1006, RCS1036, RCS1057, RCS1085, RCS1192
namespace ProtoMessage
{

    [global::ProtoBuf.ProtoContract()]
    public partial class RegisterResponse : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1, Name = @"result")]
        public Result Result { get; set; }

        [global::ProtoBuf.ProtoMember(2, Name = @"error")]
        public Error Error { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class LoginResponse : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1, Name = @"result")]
        public Result Result { get; set; }

        [global::ProtoBuf.ProtoMember(2, Name = @"error")]
        public Error Error { get; set; }

        [global::ProtoBuf.ProtoMember(3, Name = @"user")]
        public PUser User { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class CreateCharacterResponse : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1, Name = @"result")]
        public Result Result { get; set; }

        [global::ProtoBuf.ProtoMember(2, Name = @"error")]
        public Error Error { get; set; }

        [global::ProtoBuf.ProtoMember(3, Name = @"characters")]
        public global::System.Collections.Generic.List<PCharacter> Characters { get; } = new global::System.Collections.Generic.List<PCharacter>();

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class DeleteCharacterResponse : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1, Name = @"result")]
        public Result Result { get; set; }

        [global::ProtoBuf.ProtoMember(2, Name = @"error")]
        public Error Error { get; set; }

        [global::ProtoBuf.ProtoMember(3, Name = @"characters")]
        public global::System.Collections.Generic.List<PCharacter> Characters { get; } = new global::System.Collections.Generic.List<PCharacter>();

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class CharacterEnterGameResponse : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1, Name = @"result")]
        public Result Result { get; set; }

        [global::ProtoBuf.ProtoMember(2, Name = @"error")]
        public Error Error { get; set; }

        [global::ProtoBuf.ProtoMember(3, Name = @"character")]
        public PCharacter Character { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class MapCharacterEnterResponse : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1, Name = @"result")]
        public Result Result { get; set; }

        [global::ProtoBuf.ProtoMember(2, Name = @"error")]
        public Error Error { get; set; }

        [global::ProtoBuf.ProtoMember(3)]
        public int mapId { get; set; }

        [global::ProtoBuf.ProtoMember(4, Name = @"characters")]
        public global::System.Collections.Generic.List<PCharacter> Characters { get; } = new global::System.Collections.Generic.List<PCharacter>();

    }

}

#pragma warning restore CS0612, CS0618, CS1591, CS3021, IDE0079, IDE1006, RCS1036, RCS1057, RCS1085, RCS1192
#endregion
