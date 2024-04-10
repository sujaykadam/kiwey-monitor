package main

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "time"

    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
)

func clearScreen() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func bytesToGB(bytes uint64) float64 {
    return float64(bytes) / 1024 / 1024 / 1024
}

func main() {
    for {
        clearScreen()

        // CPU usage
        cpuPercent, err := cpu.Percent(0, true)
        if err == nil {
            fmt.Println("CPU Usage:")
            for i, percent := range cpuPercent {
                fmt.Printf(" CPU%d: %.2f%%\n", i, percent)
            }
        }

        // Memory usage
        vMem, err := mem.VirtualMemory()
        if err == nil {
            fmt.Printf("Memory Usage: Total: %.2f GB, Available: %.2f GB, Used: %.2f GB, UsedPercent: %.2f%%\n",
                bytesToGB(vMem.Total), bytesToGB(vMem.Available), bytesToGB(vMem.Used), vMem.UsedPercent)
        }

        // Disk usage
        diskUsage, err := disk.Usage("\\")
        if err == nil {
            fmt.Printf("Disk Usage: Total: %.2f GB, Free: %.2f GB, Used: %.2f GB, UsedPercent: %.2f%%\n",
                bytesToGB(diskUsage.Total), bytesToGB(diskUsage.Free), bytesToGB(diskUsage.Used), diskUsage.UsedPercent)
        }

        time.Sleep(1 * time.Second)
    }
}
