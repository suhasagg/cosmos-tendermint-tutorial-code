// Define a handler function for incoming IBC packets
func HandlePacket(ctx sdk.Context, k Keeper, packet ibc.Packet) (*sdk.Result, error) {
    // Retrieve the appropriate module based on the receiving port
    module, ok := k.GetModule(packet.DestinationPort)
    if !ok {
        return nil, sdkerrors.Wrap(ibc.ErrInvalidPort, packet.DestinationPort)
    }

    // Decode the packet data using the module's decoder
    data, err := module.DecodePacketData(packet.Data)
    if err != nil {
        return nil, sdkerrors.Wrap(ibc.ErrInvalidPacket, err.Error())
    }

    // Handle the packet based on the receiving channel
    switch packet.DestinationChannel {
    case "channel-1":
        // Handle packet for channel-1
        result, err := handleChannel1Packet(ctx, module, data)
        if err != nil {
            return nil, sdkerrors.Wrap(ibc.ErrPacketHandling, err.Error())
        }
        return result, nil
    case "channel-2":
        // Handle packet for channel-2
        result, err := handleChannel2Packet(ctx, module, data)
        if err != nil {
            return nil, sdkerrors.Wrap(ibc.ErrPacketHandling, err.Error())
        }
        return result, nil
    default:
        return nil, sdkerrors.Wrap(ibc.ErrInvalidChannel, packet.DestinationChannel)
    }
}

// Handler function for channel-1 packets
func handleChannel1Packet(ctx sdk.Context, module Module, data PacketData) (*sdk.Result, error) {
    // Check if the packet is valid based on the module's custom logic
    if err := module.ValidateChannel1Packet(ctx, data); err != nil {
        return nil, err
    }

    // Handle the packet according to the module's custom logic
    result, err := module.HandleChannel1Packet(ctx, data)
    if err != nil {
        return nil, err
    }

    // Return a success response
    return &sdk.Result{
        Data:   result.Data,
        Events: result.Events,
    }, nil
}
