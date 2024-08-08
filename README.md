# Raccoon Ping

This is a simple WebSocket server, to check the latency between client and server.

Sample Python code to find out the ping:
```python
import asyncio
import websockets
import time

async def ping_websocket(uri):
    try:
        async with websockets.connect(uri) as websocket:
            latencies = []

            for i in range(1, 10):
                start_time = time.time()
                await websocket.send(str(i))
                response = await websocket.recv()
                end_time = time.time()
                latency = (end_time - start_time) * 1000
                latencies.append(latency)
                print(f"Sent: {i}, Response: {response}, Latency: {latency:.2f} ms")

            average_latency = sum(latencies) / len(latencies)
            print(f"Average Latency: {average_latency:.2f} ms")

    except Exception as e:
        print(f"Failed to connect: {e}")

uri = "ws://localhost:8053/ws"

asyncio.get_event_loop().run_until_complete(ping_websocket(uri))
```

# Build and run
* Configure `LISTEN_PORT` in `.env` file

If you are using Windows to build the server for Linux, use `build.bat` script.

Or you can build it by your self using `go build -o raccoon-ping main.go`

# Requirements
- Golang >= 1.22
