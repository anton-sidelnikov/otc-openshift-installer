package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody 
type ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody instantiates a new ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody()(*ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) {
    m := &ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAttachmentItem gets the attachmentItem property value. The AttachmentItem property
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentItem()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable) {
    val, err := m.GetBackingStore().Get("attachmentItem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttachmentItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentItem(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachmentItem", m.GetAttachmentItem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAttachmentItem sets the attachmentItem property value. The AttachmentItem property
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentItem(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable)() {
    err := m.GetBackingStore().Set("attachmentItem", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBodyable 
type ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsCreateUploadSessionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachmentItem()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    SetAttachmentItem(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentItemable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
}
