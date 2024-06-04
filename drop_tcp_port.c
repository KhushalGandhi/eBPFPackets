#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <bpf/bpf_helpers.h>

// Define a BPF map to store the port number
struct bpf_map_def SEC("maps") drop_port = {
    .type = BPF_MAP_TYPE_ARRAY,
    .key_size = sizeof(__u32),
    .value_size = sizeof(__u16),
    .max_entries = 1,
};

SEC("xdp")
int drop_tcp_port(struct xdp_md *ctx) {
    // Get the IP header
    struct ethhdr *eth = bpf_hdr_pointer(ctx, 0, sizeof(*eth));
    if (!eth || eth->h_proto != bpf_htons(ETH_P_IP)) {
        return XDP_PASS;
    }

    struct iphdr *ip = (void *)eth + sizeof(*eth);
    if (!ip || ip->protocol != IPPROTO_TCP) {
        return XDP_PASS;
    }

    struct tcphdr *tcp = (void *)ip + sizeof(*ip);
    if (!tcp) {
        return XDP_PASS;
    }

    // Get the port to drop from the BPF map
    __u32 key = 0;
    __u16 *port = bpf_map_lookup_elem(&drop_port, &key);
    if (port && tcp->dest == bpf_htons(*port)) {
        // Drop the packet
        return XDP_DROP;
    }

    return XDP_PASS;
}

char __license[] SEC("license") = "GPL";
