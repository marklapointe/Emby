# Component: Emby.Server.Implementations — SharpCifs (Embedded SMB/CIFS)

**Path:** \`Emby.Server.Implementations/IO/SharpCifs/\`
**Type:** Directory | Embedded Library
**Language:** C#
**Maps to:** \`.discovery/169-emby-server-impl-sharpcifs.md\`

## Description

Embedded SharpCifs SMB/CIFS client library for network file access. Provides SMB protocol implementation, DCERPC, NetBIOS, NTLMSSP authentication, and utility classes.

## Files

### IO/SharpCifs

- `Config.cs` — Emby.Server.Implementations/IO/SharpCifs/Config.cs
- `UniAddress.cs` — Emby.Server.Implementations/IO/SharpCifs/UniAddress.cs

### Dcerpc

- `DcerpcBind.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcBind.cs
- `DcerpcBinding.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcBinding.cs
- `DcerpcConstants.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcConstants.cs
- `DcerpcError.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcError.cs
- `DcerpcException.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcException.cs
- `DcerpcHandle.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcHandle.cs
- `DcerpcMessage.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcMessage.cs
- `DcerpcPipeHandle.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcPipeHandle.cs
- `DcerpcSecurityProvider.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/DcerpcSecurityProvider.cs
- `Rpc.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Rpc.cs
- `UnicodeString.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/UnicodeString.cs
- `UUID.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/UUID.cs

### Dcerpc/Msrpc

- `LsaPolicyHandle.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/LsaPolicyHandle.cs
- `Lsarpc.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/Lsarpc.cs
- `LsarSidArrayX.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/LsarSidArrayX.cs
- `MsrpcDfsRootEnum.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcDfsRootEnum.cs
- `MsrpcEnumerateAliasesInDomain.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcEnumerateAliasesInDomain.cs
- `MsrpcGetMembersInAlias.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcGetMembersInAlias.cs
- `MsrpcLookupSids.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcLookupSids.cs
- `MsrpcLsarOpenPolicy2.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcLsarOpenPolicy2.cs
- `MsrpcQueryInformationPolicy.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcQueryInformationPolicy.cs
- `MsrpcSamrConnect2.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcSamrConnect2.cs
- `MsrpcSamrConnect4.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcSamrConnect4.cs
- `MsrpcSamrOpenAlias.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcSamrOpenAlias.cs
- `MsrpcSamrOpenDomain.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcSamrOpenDomain.cs
- `MsrpcShareEnum.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcShareEnum.cs
- `MsrpcShareGetInfo.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/MsrpcShareGetInfo.cs
- `Netdfs.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/Netdfs.cs
- `SamrAliasHandle.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/SamrAliasHandle.cs
- `Samr.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/Samr.cs
- `SamrDomainHandle.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/SamrDomainHandle.cs
- `SamrPolicyHandle.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/SamrPolicyHandle.cs
- `Srvsvc.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Msrpc/Srvsvc.cs

### Dcerpc/Ndr

- `NdrBuffer.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Ndr/NdrBuffer.cs
- `NdrException.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Ndr/NdrException.cs
- `NdrHyper.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Ndr/NdrHyper.cs
- `NdrLong.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Ndr/NdrLong.cs
- `NdrObject.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Ndr/NdrObject.cs
- `NdrShort.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Ndr/NdrShort.cs
- `NdrSmall.cs` — Emby.Server.Implementations/IO/SharpCifs/Dcerpc/Ndr/NdrSmall.cs

### Netbios

- `Lmhosts.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/Lmhosts.cs
- `Name.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/Name.cs
- `NameQueryRequest.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/NameQueryRequest.cs
- `NameQueryResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/NameQueryResponse.cs
- `NameServiceClient.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/NameServiceClient.cs
- `NameServicePacket.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/NameServicePacket.cs
- `NbtAddress.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/NbtAddress.cs
- `NbtException.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/NbtException.cs
- `NodeStatusRequest.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/NodeStatusRequest.cs
- `NodeStatusResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/NodeStatusResponse.cs
- `SessionRequestPacket.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/SessionRequestPacket.cs
- `SessionRetargetResponsePacket.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/SessionRetargetResponsePacket.cs
- `SessionServicePacket.cs` — Emby.Server.Implementations/IO/SharpCifs/Netbios/SessionServicePacket.cs

### Ntlmssp

- `NtlmFlags.cs` — Emby.Server.Implementations/IO/SharpCifs/Ntlmssp/NtlmFlags.cs
- `NtlmMessage.cs` — Emby.Server.Implementations/IO/SharpCifs/Ntlmssp/NtlmMessage.cs
- `Type1Message.cs` — Emby.Server.Implementations/IO/SharpCifs/Ntlmssp/Type1Message.cs
- `Type2Message.cs` — Emby.Server.Implementations/IO/SharpCifs/Ntlmssp/Type2Message.cs
- `Type3Message.cs` — Emby.Server.Implementations/IO/SharpCifs/Ntlmssp/Type3Message.cs

### Smb

- `ACE.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/ACE.cs
- `AllocInfo.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/AllocInfo.cs
- `AndXServerMessageBlock.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/AndXServerMessageBlock.cs
- `BufferCache.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/BufferCache.cs
- `Dfs.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Dfs.cs
- `DfsReferral.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/DfsReferral.cs
- `DosError.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/DosError.cs
- `DosFileFilter.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/DosFileFilter.cs
- `FileEntry.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/FileEntry.cs
- `IInfo.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/IInfo.cs
- `NetServerEnum2.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NetServerEnum2.cs
- `NetServerEnum2Response.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NetServerEnum2Response.cs
- `NetShareEnum.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NetShareEnum.cs
- `NetShareEnumResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NetShareEnumResponse.cs
- `NtlmAuthenticator.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NtlmAuthenticator.cs
- `NtlmChallenge.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NtlmChallenge.cs
- `NtlmContext.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NtlmContext.cs
- `NtlmPasswordAuthentication.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NtlmPasswordAuthentication.cs
- `NtStatus.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NtStatus.cs
- `NtTransQuerySecurityDesc.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NtTransQuerySecurityDesc.cs
- `NtTransQuerySecurityDescResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/NtTransQuerySecurityDescResponse.cs
- `Principal.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Principal.cs
- `SecurityDescriptor.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SecurityDescriptor.cs
- `ServerMessageBlock.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/ServerMessageBlock.cs
- `SID.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SID.cs
- `SigningDigest.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SigningDigest.cs
- `SmbAuthException.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbAuthException.cs
- `SmbComBlankResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComBlankResponse.cs
- `SmbComClose.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComClose.cs
- `SmbComCreateDirectory.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComCreateDirectory.cs
- `SmbComDelete.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComDelete.cs
- `SmbComDeleteDirectory.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComDeleteDirectory.cs
- `SmbComFindClose2.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComFindClose2.cs
- `SmbComLogoffAndX.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComLogoffAndX.cs
- `SmbComNegotiate.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComNegotiate.cs
- `SmbComNegotiateResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComNegotiateResponse.cs
- `SmbComNTCreateAndX.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComNTCreateAndX.cs
- `SmbComNTCreateAndXResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComNTCreateAndXResponse.cs
- `SmbComNtTransaction.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComNtTransaction.cs
- `SmbComNtTransactionResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComNtTransactionResponse.cs
- `SmbComOpenAndX.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComOpenAndX.cs
- `SmbComOpenAndXResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComOpenAndXResponse.cs
- `SmbComQueryInformation.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComQueryInformation.cs
- `SmbComQueryInformationResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComQueryInformationResponse.cs
- `SmbComReadAndX.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComReadAndX.cs
- `SmbComReadAndXResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComReadAndXResponse.cs
- `SmbComRename.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComRename.cs
- `SmbComSessionSetupAndX.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComSessionSetupAndX.cs
- `SmbComSessionSetupAndXResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComSessionSetupAndXResponse.cs
- `SmbComTransaction.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComTransaction.cs
- `SmbComTransactionResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComTransactionResponse.cs
- `SmbComTreeConnectAndX.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComTreeConnectAndX.cs
- `SmbComTreeConnectAndXResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComTreeConnectAndXResponse.cs
- `SmbComTreeDisconnect.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComTreeDisconnect.cs
- `SmbComWriteAndX.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComWriteAndX.cs
- `SmbComWriteAndXResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComWriteAndXResponse.cs
- `SmbComWrite.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComWrite.cs
- `SmbComWriteResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbComWriteResponse.cs
- `SmbConstants.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbConstants.cs
- `SmbException.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbException.cs
- `SmbFile.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbFile.cs
- `SmbFileExtensions.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbFileExtensions.cs
- `SmbFileFilter.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbFileFilter.cs
- `SmbFileInputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbFileInputStream.cs
- `SmbFilenameFilter.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbFilenameFilter.cs
- `SmbFileOutputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbFileOutputStream.cs
- `SmbNamedPipe.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbNamedPipe.cs
- `SmbRandomAccessFile.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbRandomAccessFile.cs
- `SmbSession.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbSession.cs
- `SmbShareInfo.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbShareInfo.cs
- `SmbTransport.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbTransport.cs
- `SmbTree.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/SmbTree.cs
- `Trans2FindFirst2.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2FindFirst2.cs
- `Trans2FindFirst2Response.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2FindFirst2Response.cs
- `Trans2FindNext2.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2FindNext2.cs
- `Trans2GetDfsReferral.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2GetDfsReferral.cs
- `Trans2GetDfsReferralResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2GetDfsReferralResponse.cs
- `Trans2QueryFSInformation.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2QueryFSInformation.cs
- `Trans2QueryFSInformationResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2QueryFSInformationResponse.cs
- `Trans2QueryPathInformation.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2QueryPathInformation.cs
- `Trans2QueryPathInformationResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2QueryPathInformationResponse.cs
- `Trans2SetFileInformation.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2SetFileInformation.cs
- `Trans2SetFileInformationResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/Trans2SetFileInformationResponse.cs
- `TransactNamedPipeInputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransactNamedPipeInputStream.cs
- `TransactNamedPipeOutputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransactNamedPipeOutputStream.cs
- `TransCallNamedPipe.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransCallNamedPipe.cs
- `TransCallNamedPipeResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransCallNamedPipeResponse.cs
- `TransPeekNamedPipe.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransPeekNamedPipe.cs
- `TransPeekNamedPipeResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransPeekNamedPipeResponse.cs
- `TransTransactNamedPipe.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransTransactNamedPipe.cs
- `TransTransactNamedPipeResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransTransactNamedPipeResponse.cs
- `TransWaitNamedPipe.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransWaitNamedPipe.cs
- `TransWaitNamedPipeResponse.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/TransWaitNamedPipeResponse.cs
- `WinError.cs` — Emby.Server.Implementations/IO/SharpCifs/Smb/WinError.cs

### Util

- `Base64.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Base64.cs
- `DES.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/DES.cs
- `Encdec.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Encdec.cs
- `Hexdump.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Hexdump.cs
- `HMACT64.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/HMACT64.cs
- `LogStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/LogStream.cs
- `MD4.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/MD4.cs
- `RC4.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/RC4.cs

### Util/Sharpen

- `AbstractMap.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/AbstractMap.cs
- `Arrays.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Arrays.cs
- `BufferedReader.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/BufferedReader.cs
- `BufferedWriter.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/BufferedWriter.cs
- `CharBuffer.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/CharBuffer.cs
- `CharSequence.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/CharSequence.cs
- `Collections.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Collections.cs
- `ConcurrentHashMap.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/ConcurrentHashMap.cs
- `DateFormat.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/DateFormat.cs
- `EnumeratorWrapper.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/EnumeratorWrapper.cs
- `Exceptions.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Exceptions.cs
- `Extensions.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Extensions.cs
- `FileInputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/FileInputStream.cs
- `FileOutputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/FileOutputStream.cs
- `FilePath.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/FilePath.cs
- `FileReader.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/FileReader.cs
- `FileWriter.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/FileWriter.cs
- `FilterInputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/FilterInputStream.cs
- `FilterOutputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/FilterOutputStream.cs
- `Hashtable.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Hashtable.cs
- `HttpURLConnection.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/HttpURLConnection.cs
- `ICallable.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/ICallable.cs
- `IConcurrentMap.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/IConcurrentMap.cs
- `IExecutor.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/IExecutor.cs
- `IFilenameFilter.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/IFilenameFilter.cs
- `IFuture.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/IFuture.cs
- `InputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/InputStream.cs
- `InputStreamReader.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/InputStreamReader.cs
- `IPrivilegedAction.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/IPrivilegedAction.cs
- `IRunnable.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/IRunnable.cs
- `Iterator.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Iterator.cs
- `LinkageError.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/LinkageError.cs
- `Matcher.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Matcher.cs
- `MD5.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/MD5.cs
- `MD5Managed.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/MD5Managed.cs
- `MessageDigest.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/MessageDigest.cs
- `NetworkStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/NetworkStream.cs
- `ObjectInputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/ObjectInputStream.cs
- `ObjectOutputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/ObjectOutputStream.cs
- `OutputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/OutputStream.cs
- `OutputStreamWriter.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/OutputStreamWriter.cs
- `PipedInputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/PipedInputStream.cs
- `PipedOutputStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/PipedOutputStream.cs
- `PrintWriter.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/PrintWriter.cs
- `Properties.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Properties.cs
- `RandomAccessFile.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/RandomAccessFile.cs
- `ReentrantLock.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/ReentrantLock.cs
- `Reference.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Reference.cs
- `Runtime.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Runtime.cs
- `SimpleDateFormat.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/SimpleDateFormat.cs
- `SocketEx.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/SocketEx.cs
- `StringTokenizer.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/StringTokenizer.cs
- `SynchronizedList.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/SynchronizedList.cs
- `Thread.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/Thread.cs
- `ThreadFactory.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/ThreadFactory.cs
- `ThreadPoolExecutor.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/ThreadPoolExecutor.cs
- `WrappedSystemStream.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Sharpen/WrappedSystemStream.cs

### Util/Transport

- `Request.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Transport/Request.cs
- `Response.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Transport/Response.cs
- `Transport.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Transport/Transport.cs
- `TransportException.cs` — Emby.Server.Implementations/IO/SharpCifs/Util/Transport/TransportException.cs

