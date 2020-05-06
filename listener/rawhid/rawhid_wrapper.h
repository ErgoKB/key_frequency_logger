#ifndef rawhid_wrapper_included_h__
#define rawhid_wrapper_included_h__


extern char returnBuf[1024];
void hid_start(void);
int hid_read(void);
void hid_close(void);

#endif
