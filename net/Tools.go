package net

import n "net"

func CheckPort(port string) error {
    var err error
    
    tcpAddress, err := n.ResolveTCPAddr("tcp4", ":" + port)
    if err != nil {
        return err
    }
    
    for i := 0; i < 3; i++ {
        listener, err := n.ListenTCP("tcp", tcpAddress)
        if err != nil {
            time.Sleep(time.Duration(100) * time.Millisecond)
            if i == 3 {
                return err
            }
            continue
        } else {
            listener.Close()
            break
        }
    }

    return nil
}