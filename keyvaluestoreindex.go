type MyItemOwnerIndexer struct {}

func (idx MyItemOwnerIndexer) Key(ctx sdk.Context, key []byte, value []byte) []byte {
    // Parse the MyItem object from the value
    var item MyItem
    err := json.Unmarshal(value, &item)
    if err != nil {
        panic(err)
    }
    // Return the owner field as the key for the index
    return []byte(item.Owner)
}

func (idx MyItemOwnerIndexer) Name() string {
    return "myitem_owner_index"
}


func NewMyApp() *MyApp {
    // Create a new instance of the Cosmos SDK application
    app := &MyApp{}
    // Create a new key-value store with the index
    kvStoreKey := sdk.NewKVStoreKey("myapp")
    kvStore := prefix.NewStore(kvStoreKey, []byte{})
    index := types.NewKVStoreIndex("myitem_owner_index", kvStore, MyItemOwnerIndexer{})
    // Register the index with the key-value store
    app.RegisterKVStoreIndex(types.StoreKey, index)
    return app
}


func (app *MyApp) GetMyItemsByOwner(ctx sdk.Context, owner string) []MyItem {
    // Get the index from the key-value store
    index := app.GetKVStoreIndex(types.StoreKey, "myitem_owner_index")
    // Get the iterator for the index with the specified owner field
    iterator := index.Iterator(ctx, []byte(owner))
    defer iterator.Close()
    // Iterate over all MyItem objects with the specified owner field
    var items []MyItem
    for ; iterator.Valid(); iterator.Next() {
        var item MyItem
        err := json.Unmarshal(iterator.Value(), &item)
        if err != nil {
            panic(err)
        }
        items = append(items, item)
    }
    return items
}
