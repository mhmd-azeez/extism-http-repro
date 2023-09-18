## Rust HTTP repro

```
Mo in D:\dylibso\x\extism-http\host λ cargo run .
   Compiling rust-host v0.1.0 (D:\dylibso\x\extism-http\host)
    Finished dev [unoptimized + debuginfo] target(s) in 4.58s
     Running `target\debug\rust-host.exe .`
thread 'main' panicked at 'called `Result::unwrap()` on an `Err` value: Call failed

Caused by:
    0: error while executing at wasm backtrace:
           0: 0xd403 - <unknown>!runtime.run$1
           1: 0xc9fe - <unknown>!runtime.run$1$gowrapper
           2:  0x459 - <unknown>!tinygo_launch
           3: 0xc8d2 - <unknown>!_start
       note: using the `WASMTIME_BACKTRACE_DETAILS=1` environment variable may show more debugging information
    1: invalid handle offset: 1114152', src\main.rs:11:51
stack backtrace:
   0:     0x7ff71c15fa62 - std::backtrace_rs::backtrace::dbghelp::trace
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\..\..\backtrace\src\backtrace\dbghelp.rs:98
   1:     0x7ff71c15fa62 - std::backtrace_rs::backtrace::trace_unsynchronized
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\..\..\backtrace\src\backtrace\mod.rs:66
   2:     0x7ff71c15fa62 - std::sys_common::backtrace::_print_fmt
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\sys_common\backtrace.rs:65
   3:     0x7ff71c15fa62 - std::sys_common::backtrace::_print::impl$0::fmt
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\sys_common\backtrace.rs:44
   4:     0x7ff71c17fa2b - core::fmt::write
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\core\src\fmt\mod.rs:1254
   5:     0x7ff71c1598ca - std::io::Write::write_fmt<std::sys::windows::stdio::Stderr>
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\io\mod.rs:1698
   6:     0x7ff71c15f7ab - std::sys_common::backtrace::_print
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\sys_common\backtrace.rs:47
   7:     0x7ff71c15f7ab - std::sys_common::backtrace::print
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\sys_common\backtrace.rs:34
   8:     0x7ff71c16209a - std::panicking::default_hook::closure$1
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:269
   9:     0x7ff71c161d00 - std::panicking::default_hook
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:288
  10:     0x7ff71c162776 - std::panicking::rust_panic_with_hook
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:691
  11:     0x7ff71c16266e - std::panicking::begin_panic_handler::closure$0
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:582
  12:     0x7ff71c160429 - std::sys_common::backtrace::__rust_end_short_backtrace<std::panicking::begin_panic_handler::closure_env$0,never$>
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\sys_common\backtrace.rs:150
  13:     0x7ff71c162380 - std::panicking::begin_panic_handler
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:578
  14:     0x7ff71c203c65 - core::panicking::panic_fmt
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\core\src\panicking.rs:67
  15:     0x7ff71c204286 - core::result::unwrap_failed
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\core\src\result.rs:1687
  16:     0x7ff71ab54155 - enum2$<core::result::Result<alloc::string::String,anyhow::Error> >::unwrap<alloc::string::String,anyhow::Error>
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca\library\core\src\result.rs:1089
  17:     0x7ff71abbbf22 - rust_host::main
                               at D:\dylibso\x\extism-http\host\src\main.rs:11
  18:     0x7ff71ac1133b - core::ops::function::FnOnce::call_once<void (*)(),tuple$<> >
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca\library\core\src\ops\function.rs:250
  19:     0x7ff71ab500be - std::sys_common::backtrace::__rust_begin_short_backtrace<void (*)(),tuple$<> >       
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca\library\std\src\sys_common\backtrace.rs:134
  20:     0x7ff71ab500be - std::sys_common::backtrace::__rust_begin_short_backtrace<void (*)(),tuple$<> >       
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca\library\std\src\sys_common\backtrace.rs:134
  21:     0x7ff71abbb911 - std::rt::lang_start::closure$0<tuple$<> >
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca\library\std\src\rt.rs:166     
  22:     0x7ff71c15392e - core::ops::function::impls::impl$2::call_once
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\core\src\ops\function.rs:287
  23:     0x7ff71c15392e - std::panicking::try::do_call
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:485
  24:     0x7ff71c15392e - std::panicking::try
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:449
  25:     0x7ff71c15392e - std::panic::catch_unwind
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panic.rs:140  
  26:     0x7ff71c15392e - std::rt::lang_start_internal::closure$2
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\rt.rs:148     
  27:     0x7ff71c15392e - std::panicking::try::do_call
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:485
  28:     0x7ff71c15392e - std::panicking::try
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panicking.rs:449
  29:     0x7ff71c15392e - std::panic::catch_unwind
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\panic.rs:140
  30:     0x7ff71c15392e - std::rt::lang_start_internal
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca/library\std\src\rt.rs:148     
  31:     0x7ff71abbb8ea - std::rt::lang_start<tuple$<> >
                               at /rustc/90c541806f23a127002de5b4038be731ba1458ca\library\std\src\rt.rs:165     
  32:     0x7ff71abbc069 - main
  33:     0x7ff71c200c50 - invoke_main
                               at D:\a\_work\1\s\src\vctools\crt\vcstartup\src\startup\exe_common.inl:78        
  34:     0x7ff71c200c50 - __scrt_common_main_seh
                               at D:\a\_work\1\s\src\vctools\crt\vcstartup\src\startup\exe_common.inl:288       
  35:     0x7ffc0a4d257d - BaseThreadInitThunk
  36:     0x7ffc0bf8aa68 - RtlUserThreadStart
error: process didn't exit successfully: `target\debug\rust-host.exe .` (exit code: 101)
```