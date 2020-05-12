#include <stdio.h>
#include "rawhid.h"

rawhid_t *hid;
char buf[128];
char returnBuf[1024], *bufPtr, *sepPtr;

static void delay_ms(unsigned int msec);

void hid_start(void) {
    while (1) {
        hid = rawhid_open_only1(0, 0, 0xFF31, 0x0074);
        if (hid == NULL) {
            delay_ms(1000);
            continue;
        }
        break;
    }
    sepPtr = bufPtr = returnBuf;
}

int hid_read(void) {
    int zeroCount = 0;
    int num, count;
    if (hid == NULL) {
        return -1;
    }
    for (count = 0; count < (bufPtr - sepPtr); count++) {
        returnBuf[count] = returnBuf[count + (sepPtr - returnBuf)];
    }
    bufPtr = returnBuf + count;
    sepPtr = returnBuf;
    while (1) {
        num = rawhid_read(hid, buf, sizeof(buf), 200);
        if (num < 0) return -1;
        if (num == 0) {
            if (++zeroCount == 3) {
                return 0;
            };
            continue;
        }
        for (count = 0; count < num; count++) {
            if (buf[count]) {
                *bufPtr = buf[count];
                ++bufPtr;
                if (bufPtr - returnBuf > sizeof(returnBuf)) {
                    bufPtr = sepPtr = returnBuf;
                    return -2;
                }
                if (buf[count] == '\n') {
                    sepPtr = bufPtr;
                }
            }
        }
        if (sepPtr != returnBuf) {
            return sepPtr - returnBuf;
        }
    }
    return 0;
}

void hid_close(void) {
    if (hid == NULL) {
        return;
    }
    rawhid_close(hid);
}

#if (defined(WIN32) || defined(WINDOWS) || defined(__WINDOWS__))
#include <windows.h>
static void delay_ms(unsigned int msec)
{
    Sleep(msec);
}
#else
#include <unistd.h>
static void delay_ms(unsigned int msec)
{
    usleep(msec * 1000);
}
#endif
