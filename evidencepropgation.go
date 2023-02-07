type Evidence struct {
    Height  int64
    Address string
    Evidence []byte
}

func gossipEvidence(ev Evidence) {
    peers := getPeers()
    for _, peer := range peers {
        go sendEvidence(peer, ev)
    }
}


func sendEvidence(peer string, ev Evidence) error {
    // Connect to the peer
    conn, err := net.Dial("tcp", peer)
    if err != nil {
        return err
    }
    defer conn.Close()

    // Send the evidence
    err = binary.Write(conn, binary.BigEndian, ev)
    if err != nil {
        return err
    }

    return nil
}

func receiveEvidence() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Error listening: %v", err)
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatalf("Error accepting: %v", err)
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    var ev Evidence
    err := binary.Read(conn, binary.BigEndian, &ev)
    if err != nil {
        log.Fatalf("Error reading: %v", err)
    }

    // Add the evidence to a cache
    addEvidenceToCache(ev)
}
