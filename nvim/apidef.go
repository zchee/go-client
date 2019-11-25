// +build ignore

// This file defines the Nvim remote API using Go syntax. Run the 'go generate'
// command to convert this file to the API implementation in apiimp.go.
//
// The go generate command runs the apitool program. Run
//
//  go run apitool.go --help
//
// to learn more about the apitool program.

package main

// BufferLineCount returns the number of lines in the buffer.
func BufferLineCount(buffer Buffer) int {
	name(nvim_buf_line_count)
}

// BufferLines retrieves a line range from a buffer.
//
// Indexing is zero-based, end-exclusive. Negative indices are interpreted as
// length+1+index, i e -1 refers to the index past the end. So to get the last
// element set start=-2 and end=-1.
//
// Out-of-bounds indices are clamped to the nearest valid value, unless strict
// = true.
func BufferLines(buffer Buffer, start int, end int, strict bool) [][]byte {
	name(nvim_buf_get_lines)
}

// AttachBuffer activate updates from this buffer to the current channel.
//
// If sendBuffer is true, initial notification should contain the whole buffer.
// If false, the first notification will be a `nvim_buf_lines_event`.
// Otherwise, the first notification will be a `nvim_buf_changedtick_event`
//
// opts is optional parameters. Currently not used.
//
// returns whether the updates couldn't be enabled because the buffer isn't loaded or opts contained an invalid key.
func AttachBuffer(buffer Buffer, sendBuffer bool, opts map[string]interface{}) bool {
	name(nvim_buf_attach)
}

// DetachBuffer deactivate updates from this buffer to the current channel.
//
// returns whether the updates couldn't be disabled because the buffer isn't loaded.
func DetachBuffer(buffer Buffer) bool {
	name(nvim_buf_detach)
}

// SetBufferLines replaces a line range on a buffer.
//
// Indexing is zero-based, end-exclusive. Negative indices are interpreted as
// length+1+index, ie -1 refers to the index past the end. So to change or
// delete the last element set start=-2 and end=-1.
//
// To insert lines at a given index, set both start and end to the same index.
// To delete a range of lines, set replacement to an empty array.
//
// Out-of-bounds indices are clamped to the nearest valid value, unless strict
// = true.
func SetBufferLines(buffer Buffer, start int, end int, strict bool, replacement [][]byte) {
	name(nvim_buf_set_lines)
}

// BufferOffset returns the byte offset for a line.
//
// Line 1 (index=0) has offset 0. UTF-8 bytes are counted. EOL is one byte.
// 'fileformat' and 'fileencoding' are ignored. The line index just after the
// last line gives the total byte-count of the buffer. A final EOL byte is
// counted if it would be written, see 'eol'.
//
// Unlike |line2byte()|, throws error for out-of-bounds indexing.
// Returns -1 for unloaded buffer.
func BufferOffset(buffer Buffer, index int) int {
	name(nvim_buf_get_offset)
}

// BufferVar gets a buffer-scoped (b:) variable.
func BufferVar(buffer Buffer, name string) interface{} {
	name(nvim_buf_get_var)
}

// BufferChangedTick gets a changed tick of a buffer.
func BufferChangedTick(buffer Buffer) int {
	name(nvim_buf_get_changedtick)
}

// BufferKeymap gets a list of buffer-local mappings.
func BufferKeyMap(buffer Buffer, mode string) []*Mapping {
	name(nvim_buf_get_keymap)
}

// SetBufferKeyMap sets a buffer-local mapping for the given mode.
//
// see
//  :help nvim_set_keymap()
func SetBufferKeyMap(buffer Buffer, mode, lhs, rhs string, opts map[string]bool) {
	name(nvim_buf_set_keymap)
}

// DeleteBufferKeyMap unmaps a buffer-local mapping for the given mode.
//
// see
//  :help nvim_del_keymap()
func DeleteBufferKeyMap(buffer Buffer, mode, lhs string) {
	name(nvim_buf_del_keymap)
}

// BufferCommands gets a map of buffer-local user-commands.
//
// opts is optional parameters. Currently not used.
func BufferCommands(buffer Buffer, opts map[string]interface{}) map[string]*Command {
	name(nvim_buf_get_commands)
}

// SetBufferVar sets a buffer-scoped (b:) variable.
func SetBufferVar(buffer Buffer, name string, value interface{}) {
	name(nvim_buf_set_var)
}

// DeleteBufferVar removes a buffer-scoped (b:) variable.
func DeleteBufferVar(buffer Buffer, name string) {
	name(nvim_buf_del_var)
}

// BufferOption gets a buffer option value.
func BufferOption(buffer Buffer, name string) interface{} {
	name(nvim_buf_get_option)
}

// SetBufferOption sets a buffer option value. The value nil deletes the option
// in the case where there's a global fallback.
func SetBufferOption(buffer Buffer, name string, value interface{}) {
	name(nvim_buf_set_option)
}

// BufferNumber gets a buffer's number.
//
// Deprecated: Use int(buffer) to get the buffer's number as an integer.
func BufferNumber(buffer Buffer) int {
	name(nvim_buf_get_number)
	deprecatedSince(2)
}

// BufferName gets the full file name of a buffer.
func BufferName(buffer Buffer) string {
	name(nvim_buf_get_name)
}

// SetBufferName sets the full file name of a buffer.
// BufFilePre/BufFilePost are triggered.
func SetBufferName(buffer Buffer, name string) {
	name(nvim_buf_set_name)
}

// IsBufferLoaded Checks if a buffer is valid and loaded.
// See api-buffer for more info about unloaded buffers.
func IsBufferLoaded(buffer Buffer) bool {
	name(nvim_buf_is_loaded)
}

// IsBufferValid returns true if the buffer is valid.
func IsBufferValid(buffer Buffer) bool {
	name(nvim_buf_is_valid)
}

// BufferMark returns the (row,col) of the named mark.
func BufferMark(buffer Buffer, name string) [2]int {
	name(nvim_buf_get_mark)
}

// BufferExtmarkByID returns position for a given extmark id.
func BufferExtmarkByID(buffer Buffer, nsID int, id int) []int {
	name(nvim_buf_get_extmark_by_id)
}

// BufferExtmarks gets extmarks in "traversal order" from a charwise region defined by
// buffer positions (inclusive, 0-indexed).
//
// Region can be given as (row,col) tuples, or valid extmark ids (whose
// positions define the bounds). 0 and -1 are understood as (0,0) and (-1,-1)
// respectively, thus the following are equivalent:
//
//   BufferExtmarks(0, myNS, 0, -1, {})
//   BufferExtmarks(0, myNS, [0,0], [-1,-1], {})
//
// If `end` is less than `start`, traversal works backwards. (Useful
// with `limit`, to get the first marks prior to a given position.)
//
// The `opts` is additional options. Supports the key:
//   limit: (int) Maximum number of marks to return.
func BufferExtmarks(buffer Buffer, nsID int, start interface{}, end interface{}, opt map[string]interface{}) []interface{} {
	name(nvim_buf_get_extmarks)
}

// SetBufferExtmark creates or updates an extmark.
//
// To create a new extmark, pass id=0. The extmark id will be returned.
// To move an existing mark, pass its id.
//
// It is also allowed to create a new mark by passing in a previously unused
// id, but the caller must then keep track of existing and unused ids itself.
// (Useful over RPC, to avoid waiting for the return value.)
//
// Currently opts arg not used.
func SetBufferExtmark(buffer Buffer, nsID int, extmarkID int, line int, col int, opts map[string]interface{}) int {
	name(nvim_buf_set_extmark)
}

// DeleteBufferExtmark removes an extmark.
func DeleteBufferExtmark(buffer Buffer, nsID int, extmarkID int) bool {
	name(nvim_buf_del_extmark)
}

// AddBufferHighlight adds a highlight to buffer and returns the source id of
// the highlight.
//
// AddBufferHighlight can be used for plugins which dynamically generate
// highlights to a buffer (like a semantic highlighter or linter). The function
// adds a single highlight to a buffer. Unlike matchaddpos() highlights follow
// changes to line numbering (as lines are inserted/removed above the
// highlighted line), like signs and marks do.
//
// The srcID is useful for batch deletion/updating of a set of highlights. When
// called with srcID = 0, an unique source id is generated and returned.
// Successive calls can pass in it as srcID to add new highlights to the same
// source group. All highlights in the same group can then be cleared with
// ClearBufferHighlight. If the highlight never will be manually deleted pass
// in -1 for srcID.
//
// If hlGroup is the empty string no highlight is added, but a new srcID is
// still returned. This is useful for an external plugin to synchronously
// request an unique srcID at initialization, and later asynchronously add and
// clear highlights in response to buffer changes.
//
// The startCol and endCol parameters specify the range of columns to
// highlight. Use endCol = -1 to highlight to the end of the line.
func AddBufferHighlight(buffer Buffer, srcID int, hlGroup string, line int, startCol int, endCol int) int {
	name(nvim_buf_add_highlight)
}

// ClearBufferNamespace clears namespaced objects, highlights and virtual text, from a line range.
//
// To clear the namespace in the entire buffer, pass in 0 and -1 to
// line_start and line_end respectively.
func ClearBufferNamespace(buffer Buffer, nsID int, lineStart int, lineEnd int) {
	name(nvim_buf_clear_namespace)
}

// ClearBufferHighlight clears highlights from a given source group and a range
// of lines.
//
// To clear a source group in the entire buffer, pass in 1 and -1 to startLine
// and endLine respectively.
//
// The lineStart and lineEnd parameters specify the range of lines to clear.
// The end of range is exclusive. Specify -1 to clear to the end of the file.
//
// Deprecated: Use ClearBufferNamespace() instead.
func ClearBufferHighlight(buffer Buffer, srcID int, startLine int, endLine int) {
	name(nvim_buf_clear_highlight)
}

// SetBufferVirtualText sets the virtual text (annotation) for a buffer line.
//
// By default (and currently the only option) the text will be placed after
// the buffer text. Virtual text will never cause reflow, rather virtual
// text will be truncated at the end of the screen line. The virtual text will
// begin one cell (|lcs-eol| or space) after the ordinary text.
//
// Namespaces are used to support batch deletion/updating of virtual text.
// To create a namespace, use CreateNamespace(). Virtual text is
// cleared using ClearBufferNamespace(). The same `nsID` can be used for
// both virtual text and highlights added by AddBufferHighlight(), both
// can then be cleared with a single call to ClearBufferNamespace(). If the
// virtual text never will be cleared by an API call, pass `nsID = -1`.
//
// As a shorthand, `nsID = 0` can be used to create a new namespace for the virtual text, the allocated id is then returned.
//
// The `opts` is optional parameters. Currently not used.
//
// The returns the nsID that was used.
func SetBufferVirtualText(buffer Buffer, nsID int, line int, chunks []VirtualTextChunk, opts map[string]interface{}) int {
	name(nvim_buf_set_virtual_text)
}

// BufferVirtualText gets the virtual text (annotation) for a buffer line.
//
// The virtual text is returned as list of lists, whereas the inner lists have
// either one or two elements. The first element is the actual text, the
// optional second element is the highlight group.
//
// The format is exactly the same as given to SetBufferVirtualText.
//
// If there is no virtual text associated with the given line, an empty list
// is returned.
func BufferVirtualText(buffer Buffer, lnum int) []VirtualTextChunk {
	name(nvim_buf_get_virtual_text)
}

// TabpageWindows returns the windows in a tabpage.
func TabpageWindows(tabpage Tabpage) []Window {
	name(nvim_tabpage_list_wins)
}

// TabpageVar gets a tab-scoped (t:) variable.
func TabpageVar(tabpage Tabpage, name string) interface{} {
	name(nvim_tabpage_get_var)
}

// SetTabpageVar sets a tab-scoped (t:) variable.
func SetTabpageVar(tabpage Tabpage, name string, value interface{}) {
	name(nvim_tabpage_set_var)
}

// DeleteTabpageVar removes a tab-scoped (t:) variable.
func DeleteTabpageVar(tabpage Tabpage, name string) {
	name(nvim_tabpage_del_var)
}

// TabpageWindow gets the current window in a tab page.
func TabpageWindow(tabpage Tabpage) Window {
	name(nvim_tabpage_get_win)
}

// TabpageNumber gets the tabpage number from the tabpage handle.
func TabpageNumber(tabpage Tabpage) int {
	name(nvim_tabpage_get_number)
}

// IsTabpageValid checks if a tab page is valid.
func IsTabpageValid(tabpage Tabpage) bool {
	name(nvim_tabpage_is_valid)
}

// AttachUI registers the client as a remote UI. After this method is called,
// the client will receive redraw notifications.
//
//  :help rpc-remote-ui
//
// The redraw notification method has variadic arguments. Register a handler
// for the method like this:
//
//  v.RegisterHandler("redraw", func(updates ...[]interface{}) {
//      for _, update := range updates {
//          // handle update
//      }
//  })
func AttachUI(width int, height int, options map[string]interface{}) {
	name(nvim_ui_attach)
}

// DetachUI unregisters the client as a remote UI.
func DetachUI() {
	name(nvim_ui_detach)
}

// TryResizeUI notifies Nvim that the client window has resized. If possible,
// Nvim will send a redraw request to resize.
func TryResizeUI(width int, height int) {
	name(nvim_ui_try_resize)
}

// SetUIOption sets a UI option.
func SetUIOption(name string, value interface{}) {
	name(nvim_ui_set_option)
}

// TryResizeUIGrid tell Nvim to resize a grid. Triggers a grid_resize event with the requested
// grid size or the maximum size if it exceeds size limits.
//
// On invalid grid handle, fails with error.
func TryResizeUIGrid(grid, width, height int) {
	name(nvim_ui_try_resize_grid)
}

// SetPumHeight tells Nvim the number of elements displaying in the popumenu, to decide
// <PageUp> and <PageDown> movement.
//
// height is popupmenu height, must be greater than zero.
func SetPumHeight(height int) {
	name(nvim_ui_pum_set_height)
}

// Command executes a single ex command.
func Command(cmd string) {
	name(nvim_command)
}

// HLByID gets a highlight definition by id.
func HLByID(id int, rgb bool) *HLAttrs {
	name(nvim_get_hl_by_id)
}

// HLByName gets a highlight definition by name.
func HLByName(name string, rgb bool) *HLAttrs {
	name(nvim_get_hl_by_name)
}

// FeedKeys Pushes keys to the Nvim user input buffer. Options can be a string
// with the following character flags:
//
//  m:  Remap keys. This is default.
//  n:  Do not remap keys.
//  t:  Handle keys as if typed; otherwise they are handled as if coming from a
//     mapping. This matters for undo, opening folds, etc.
func FeedKeys(keys string, mode string, escapeCSI bool) {
	name(nvim_feedkeys)
}

// Input pushes bytes to the Nvim low level input buffer.
//
// Unlike FeedKeys, this uses the lowest level input buffer and the call is not
// deferred. It returns the number of bytes actually written(which can be less
// than what was requested if the buffer is full).
func Input(keys string) int {
	name(nvim_input)
}

// InputMouse send mouse event from GUI.
//
// The call is non-blocking. It doesn't wait on any resulting action, but
// queues the event to be processed soon by the event loop.
func InputMouse(button, action, modifier string, grid, row, col int) {
	name(nvim_input_mouse)
}

// ReplaceTermcodes replaces any terminal code strings by byte sequences. The
// returned sequences are Nvim's internal representation of keys, for example:
//
//  <esc> -> '\x1b'
//  <cr>  -> '\r'
//  <c-l> -> '\x0c'
//  <up>  -> '\x80ku'
//
// The returned sequences can be used as input to feedkeys.
func ReplaceTermcodes(str string, fromPart bool, doLT bool, special bool) string {
	name(nvim_replace_termcodes)
}

// CommandOutput executes a single ex command and returns the output.
func CommandOutput(cmd string) string {
	name(nvim_command_output)
}

// Eval evaluates the expression expr using the Vim internal expression
// evaluator.
//
//  :help expression
func Eval(expr string) interface{} {
	name(nvim_eval)
}

// StringWidth returns the number of display cells the string occupies. Tab is
// counted as one cell.
func StringWidth(s string) int {
	name(nvim_strwidth)
}

// RuntimePaths returns a list of paths contained in the runtimepath option.
func RuntimePaths() []string {
	name(nvim_list_runtime_paths)
}

// SetCurrentDirectory changes the Vim working directory.
func SetCurrentDirectory(dir string) {
	name(nvim_set_current_dir)
}

// CurrentLine gets the current line in the current buffer.
func CurrentLine() []byte {
	name(nvim_get_current_line)
}

// SetCurrentLine sets the current line in the current buffer.
func SetCurrentLine(line []byte) {
	name(nvim_set_current_line)
}

// DeleteCurrentLine deletes the current line in the current buffer.
func DeleteCurrentLine() {
	name(nvim_del_current_line)
}

// Var gets a global (g:) variable.
func Var(name string) interface{} {
	name(nvim_get_var)
}

// SetVar sets a global (g:) variable.
func SetVar(name string, value interface{}) {
	name(nvim_set_var)
}

// DeleteVar removes a global (g:) variable.
func DeleteVar(name string) {
	name(nvim_del_var)
}

// VVar gets a vim (v:) variable.
func VVar(name string) interface{} {
	name(nvim_get_vvar)
}

// SetVVar sets a v: variable, if it is not readonly.
func SetVVar(name string, value interface{}) {
	name(nvim_set_vvar)
}

// Option gets an option.
func Option(name string) interface{} {
	name(nvim_get_option)
}

// SetOption sets an option.
func SetOption(name string, value interface{}) {
	name(nvim_set_option)
}

// WriteOut writes a message to vim output buffer. The string is split and
// flushed after each newline. Incomplete lines are kept for writing later.
func WriteOut(str string) {
	name(nvim_out_write)
}

// WriteErr writes a message to vim error buffer. The string is split and
// flushed after each newline. Incomplete lines are kept for writing later.
func WriteErr(str string) {
	name(nvim_err_write)
}

// WritelnErr writes prints str and a newline as an error message.
func WritelnErr(str string) {
	name(nvim_err_writeln)
}

// Buffers returns the current list of buffers.
func Buffers() []Buffer {
	name(nvim_list_bufs)
}

// CurrentBuffer returns the current buffer.
func CurrentBuffer() Buffer {
	name(nvim_get_current_buf)
}

// SetCurrentBuffer sets the current buffer.
func SetCurrentBuffer(buffer Buffer) {
	name(nvim_set_current_buf)
}

// Windows returns the current list of windows.
func Windows() []Window {
	name(nvim_list_wins)
}

// CurrentWindow returns the current window.
func CurrentWindow() Window {
	name(nvim_get_current_win)
}

// SetCurrentWindow sets the current window.
func SetCurrentWindow(window Window) {
	name(nvim_set_current_win)
}

// CreateBuffer creates a new, empty, unnamed buffer.
func CreateBuffer(listed, scratch bool) Buffer {
	name(nvim_create_buf)
}

// OpenWindow opens a new window.
//
// Currently this is used to open floating and external windows.
// Floats are windows that are drawn above the split layout, at some anchor
// position in some other window. Floats can be drawn internally or by external
// GUI with the |ui-multigrid| extension. External windows are only supported
// with multigrid GUIs, and are displayed as separate top-level windows.
//
// For a general overview of floats, see |api-floatwin|.
//
// Exactly one of External and Relative must be specified. The Width and
// Height of the new window must be specified.
//
// With relative=editor (row=0,col=0) refers to the top-left corner of the
// screen-grid and (row=Lines-1,col=Columns-1) refers to the bottom-right
// corner. Fractional values are allowed, but the builtin implementation
// (used by non-multigrid UIs) will always round down to nearest integer.
//
// Out-of-bounds values, and configurations that make the float not fit inside
// the main editor, are allowed. The builtin implementation truncates values
// so floats are fully within the main screen grid. External GUIs
// could let floats hover outside of the main window like a tooltip, but
// this should not be used to specify arbitrary WM screen positions.
//
// The returns the window handle or 0 when error.
func OpenWindow(buffer Buffer, enter bool, config *WindowConfig) Window {
	name(nvim_open_win)
}

// Tabpages returns the current list of tabpages.
func Tabpages() []Tabpage {
	name(nvim_list_tabpages)
}

// CurrentTabpage returns the current tabpage.
func CurrentTabpage() Tabpage {
	name(nvim_get_current_tabpage)
}

// SetCurrentTabpage sets the current tabpage.
func SetCurrentTabpage(tabpage Tabpage) {
	name(nvim_set_current_tabpage)
}

// CreateNamespace creates a new namespace, or gets an existing one.
//
// Namespaces are used for buffer highlights and virtual text, see
// AddBufferHighlight() and SetBufferVirtualText().
//
// Namespaces can be named or anonymous. If `name` matches an existing
// namespace, the associated id is returned. If `name` is an empty string
// a new, anonymous namespace is created.
//
// The returns the namespace ID.
func CreateNamespace(name string) int {
	name(nvim_create_namespace)
}

// Namespaces gets existing named namespaces
//
// The return dict that maps from names to namespace ids.
func Namespaces() map[string]int {
	name(nvim_get_namespaces)
}

// Paste pastes at cursor, in any mode.
//
// Invokes the `vim.paste` handler, which handles each mode appropriately.
// Sets redo/undo. Faster than Input(). Lines break at LF ("\n").
//
// Errors ('nomodifiable', `vim.paste()` failure, …) are reflected in `err`
// but do not affect the return value (which is strictly decided by
// `vim.paste()`).  On error, subsequent calls are ignored ("drained") until
// the next paste is initiated (phase 1 or -1).
//
//  data
// multiline input. May be binary (containing NUL bytes).
//  crlf
// also break lines at CR and CRLF.
//  phase
// -1 is paste in a single call (i.e. without streaming).
//
// To "stream" a paste, call Paste sequentially with these `phase` values:
//  1: starts the paste (exactly once)
//  2: continues the paste (zero or more times)
//  3: ends the paste (exactly once)
func Paste(data string, crlf bool, phase int) bool {
	name(nvim_paste)
}

// Put puts text at cursor, in any mode.
//
// Compare :put and p which are always linewise.
//
// lines is readfile() style list of lines.
//
// type is edit behavior: any getregtype() result, or:
//   "b": blockwise-visual mode (may include width, e.g. "b3")
//   "c": characterwise mode
//   "l": linewise mode
//   "" : guess by contents, see setreg()
// after is insert after cursor (like `p`), or before (like `P`).
//
// follow is place cursor at end of inserted text.
func Put(lines []string, typ string, after bool, follow bool) {
	name(nvim_put)
}

// Subscribe subscribes to a Nvim event.
func Subscribe(event string) {
	name(nvim_subscribe)
}

// Unsubscribe unsubscribes to a Nvim event.
func Unsubscribe(event string) {
	name(nvim_unsubscribe)
}

func ColorByName(name string) int {
	name(nvim_get_color_by_name)
}

func ColorMap() map[string]int {
	name(nvim_get_color_map)
}

// Context gets a map of the current editor state.
//
// The `opts` is optional parameters.
//  types: List of context-types ("regs", "jumps", "bufs", "gvars" or etc) to gather, or empty for "all".
func Context(opts map[string][]string) map[string]interface{} {
	name(nvim_get_context)
}

// LoadContext sets the current editor state from the given context map.
func LoadContext(dict map[string]interface{}) interface{} {
	name(nvim_load_context)
}

// Mode gets Nvim's current mode.
func Mode() Mode {
	name(nvim_get_mode)
	returnPtr()
}

func KeyMap(mode string) []*Mapping {
	name(nvim_get_keymap)
}

// SetKeyMap sets a global |mapping| for the given mode.
//
// To set a buffer-local mapping, use SetBufferKeyMap().
//
// Unlike :map, leading/trailing whitespace is accepted as part of the {lhs}
// or {rhs}.
// Empty {rhs} is <Nop>. keycodes are replaced as usual.
//
//  mode
// mode short-name (map command prefix: "n", "i", "v", "x", …) or "!" for :map!, or empty string for :map.
//
//  lhs
// Left-hand-side {lhs} of the mapping.
//
//  rhs
// Right-hand-side {rhs} of the mapping.
//
//  opts
// Optional parameters map. Accepts all :map-arguments as keys excluding <buffer> but including noremap.
// Values are Booleans. Unknown key is an error.
func SetKeyMap(mode, lhs, rhs string, opts map[string]bool) {
	name(nvim_set_keymap)
}

// DeleteKeyMap unmaps a global mapping for the given mode.
//
// To unmap a buffer-local mapping, use DeleteBufferKeyMap().
//
// see
//  :help nvim_set_keymap()
func DeleteKeyMap(mode, lhs string) {
	name(nvim_del_keymap)
}

// Commands gets a map of global (non-buffer-local) Ex commands.
// Currently only user-commands are supported, not builtin Ex commands.
//
// opts is optional parameters. Currently only supports {"builtin":false}.
func Commands(opts map[string]interface{}) map[string]*Command {
	name(nvim_get_commands)
}

func APIInfo() []interface{} {
	name(nvim_get_api_info)
}

// SetClientInfo identify the client for nvim.
//
// Can be called more than once, but subsequent calls will remove earlier info, which should be resent if it is still valid.
// (This could happen if a library first identifies the channel, and a plugin using that library later overrides that info)
func SetClientInfo(name string, version *ClientVersion, typ string, methods map[string]*ClientMethod, attributes ClientAttributes) {
	name(nvim_set_client_info)
}

// ChannelInfo get information about a channel.
func ChannelInfo(channel int) *Channel {
	name(nvim_get_chan_info)
}

// Channels get information about all open channels.
func Channels() []*Channel {
	name(nvim_list_chans)
}

// ParseExpression parse a VimL expression.
func ParseExpression(expr string, flags string, highlight bool) map[string]interface{} {
	name(nvim_parse_expression)
}

// UIs gets a list of dictionaries representing attached UIs.
func UIs() []*UI {
	name(nvim_list_uis)
}

// ProcChildren gets the immediate children of process `pid`.
func ProcChildren(pid int) []*Process {
	name(nvim_get_proc_children)
}

// Proc gets info describing process `pid`.
func Proc(pid int) Process {
	name(nvim_get_proc)
}

// SelectPopupmenuItem selects an item in the completion popupmenu.
//
// If |ins-completion| is not active this API call is silently ignored.
// Useful for an external UI using |ui-popupmenu| to control the popupmenu
// with the mouse. Can also be used in a mapping; use <cmd> |:map-cmd| to
// ensure the mapping doesn't end completion mode.
//
// The `opts` optional parameters. Reserved for future use.
func SelectPopupmenuItem(item int, insert, finish bool, opts map[string]interface{}) {
	name(nvim_select_popupmenu_item)
}

// WindowBuffer returns the current buffer in a window.
func WindowBuffer(window Window) Buffer {
	name(nvim_win_get_buf)
}

// SetBufferToWindow sets the current buffer in a window, without side-effects.
func SetBufferToWindow(window Window, buffer Buffer) {
	name(nvim_win_set_buf)
}

// WindowCursor returns the cursor position in the window.
func WindowCursor(window Window) [2]int {
	name(nvim_win_get_cursor)
}

// SetWindowCursor sets the cursor position in the window to the given position.
func SetWindowCursor(window Window, pos [2]int) {
	name(nvim_win_set_cursor)
}

// WindowHeight returns the window height.
func WindowHeight(window Window) int {
	name(nvim_win_get_height)
}

// SetWindowHeight sets the window height.
func SetWindowHeight(window Window, height int) {
	name(nvim_win_set_height)
}

// WindowWidth returns the window width.
func WindowWidth(window Window) int {
	name(nvim_win_get_width)
}

// SetWindowWidth sets the window width.
func SetWindowWidth(window Window, width int) {
	name(nvim_win_set_width)
}

// WindowVar gets a window-scoped (w:) variable.
func WindowVar(window Window, name string) interface{} {
	name(nvim_win_get_var)
}

// SetWindowVar sets a window-scoped (w:) variable.
func SetWindowVar(window Window, name string, value interface{}) {
	name(nvim_win_set_var)
}

// DeleteWindowVar removes a window-scoped (w:) variable.
func DeleteWindowVar(window Window, name string) {
	name(nvim_win_del_var)
}

// WindowOption gets a window option.
func WindowOption(window Window, name string) interface{} {
	name(nvim_win_get_option)
}

// SetWindowOption sets a window option.
func SetWindowOption(window Window, name string, value interface{}) {
	name(nvim_win_set_option)
}

// WindowPosition gets the window position in display cells. First position is zero.
func WindowPosition(window Window) [2]int {
	name(nvim_win_get_position)
}

// WindowTabpage gets the tab page that contains the window.
func WindowTabpage(window Window) Tabpage {
	name(nvim_win_get_tabpage)
}

// WindowNumber gets the window number from the window handle.
func WindowNumber(window Window) int {
	name(nvim_win_get_number)
}

// IsWindowValid returns true if the window is valid.
func IsWindowValid(window Window) bool {
	name(nvim_win_is_valid)
}

// SetWindowConfig configure window position. Currently this is only used to configure
// floating and external windows (including changing a split window to these
// types).
//
// See documentation at |nvim_open_win()|, for the meaning of parameters.
//
// When reconfiguring a floating window, absent option keys will not be
// changed. The following restriction apply: `row`, `col` and `relative`
// must be reconfigured together. Only changing a subset of these is an error.
func SetWindowConfig(window Window, config map[string]interface{}) {
	name(nvim_win_set_config)
}

// WindowConfig return window configuration.
//
// Return a dictionary containing the same config that can be given to
// |nvim_open_win()|.
//
// `relative` will be an empty string for normal windows.
func WindowConfig(window Window) map[string]interface{} {
	name(nvim_win_get_config)
}

// CloseWindow close a window.
//
// This is equivalent to |:close| with count except that it takes a window id.
func CloseWindow(window Window, force bool) {
	name(nvim_win_close)
}
