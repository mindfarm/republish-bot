CREATE TABLE releases(
    id INTEGER,
    title TEXT,
    content TEXT,
    link TEXT
);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.5', '<h2>gopls/v0.5.5</h2>
<p>This is a patch release to fix two bugs in <code>gopls/v0.5.4</code>.</p>
<h2>Fixes</h2>
<h3>Excessive reloading of packages outside of GOPATH or a module</h3>
<p>See <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="749915462" data-permission-text="Title is private" data-url="https://github.com/golang/go/issues/42813" data-hovercard-type="issue" data-hovercard-url="/golang/go/issues/42813/hovercard" href="https://github.com/golang/go/issues/42813">golang/go#42813</a>.</p>
<h3>File corruption with CRLF line endings and <code>//</code>-style comments</h3>
<p><a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="744270596" data-permission-text="Title is private" data-url="https://github.com/golang/go/issues/42646" data-hovercard-type="issue" data-hovercard-url="/golang/go/issues/42646/hovercard" href="https://github.com/golang/go/issues/42646">golang/go#42646</a> was supposed to have been fixed in <code>gopls/v0.5.4</code>, but it was not. <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="754412985" data-permission-text="Title is private" data-url="https://github.com/golang/go/issues/42923" data-hovercard-type="issue" data-hovercard-url="/golang/go/issues/42923/hovercard" href="https://github.com/golang/go/issues/42923">golang/go#42923</a> was reported and fixed.</p>
<hr>
<p>A full list of all issues fixed can be found in the <a href="https://github.com/golang/go/milestone/189?closed=1">gopls/v0.5.5 milestone</a>. To report a new problem, please file a new issue at <a href="https://golang.org/issues/new" rel="nofollow">https://golang.org/issues/new</a>.</p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.5', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.5-pre.1', '<p>gopls/v0.5.5-pre.1</p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.5-pre.1', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.4', '<h2>Features</h2>
<h3>Opening a project that contains a module in a subdirectory</h3>
<p>Previously, <code>gopls</code> required that you open your editor exactly at or below the module root (the directory containing the <code>go.mod</code>). Now, you can open a directory that contains <strong>exactly one</strong> module in a subdirectory, and <code>gopls</code> will work as expected. For details on multi-module workspaces, see below.</p>
<h3>Removal of the granular <code>go.mod</code> upgrade codelenses</h3>
<p>Previously, we offered a code lens to suggest upgrades for each <code>require</code> in a <code>go.mod</code> file. In anticipation of changes that limit the amount that <code>gopls</code> accesses the network, we have decided to remove and reevaluate this feature. Users had mentioned that the code lenses cluttered their <code>go.mod</code> files, especially if they didn''t actually want to upgrade. <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="597596353" data-permission-text="Title is private" data-url="https://github.com/golang/go/issues/38339" data-hovercard-type="issue" data-hovercard-url="/golang/go/issues/38339/hovercard" href="https://github.com/golang/go/issues/38339">golang/go#38339</a> tracks the work to revamp this feature. An "Upgrade all dependencies" code lens should still appear at the top of your <code>go.mod</code> file.</p>
<h3>Improved error message reports</h3>
<p>Previously, critical error messages were reported as message pop-up that would re-trigger as you type. Many users would find this annoying. We have changed the approach to show error messages as progress reports, which should be less intrusive and appear more like status bars.</p>
<h3>Improved memory usage for workspaces with multiple folders</h3>
<p>We are now using a coarser cache key for package type information. If you use the gopls daemon, this may reduce your total memory usage.</p>
<h3>Experimental</h3>
<h4>Multi-module workspace support</h4>
<p>The proposal described in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="451231859" data-permission-text="Title is private" data-url="https://github.com/golang/go/issues/32394" data-hovercard-type="issue" data-hovercard-url="/golang/go/issues/32394/hovercard" href="https://github.com/golang/go/issues/32394">golang/go#32394</a> is still in development and off by default. See our progress by tracking the multi-module workspace milestone and project.</p>
<p>Enable multi-module workspace support by adding the following to your settings:</p>
<div class="highlight highlight-source-js"><pre><span class="pl-s">"gopls"</span>: <span class="pl-kos">{</span>
    <span class="pl-s">"experimentalWorkspaceModule"</span>: <span class="pl-c1">true</span><span class="pl-kos">,</span>
<span class="pl-kos">}</span></pre></div>
<p>With this setting, you will be able to open a directory that contains multiple modules or a directory that contains nested modules.</p>
<p>Give this a try if you''re interested in this new feature, but please note that it is still very experimental. Please file issues if you encounter bugs.</p>
<h2>Fixes</h2>
<h3>File corruption with CRLF line endings and <code>/**/</code>-style comments</h3>
<p>Previously, when you organized the imports in a file with CRLF line endings and multi-line comments, the formatter might output incorrect content for the file, rendering it invalid Go code. This issue has popped up a number of times, but we believe it has finally been fixed for good. If you are using Windows with CRLF line ending, please report any regressions. For full details, see <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="744270596" data-permission-text="Title is private" data-url="https://github.com/golang/go/issues/42646" data-hovercard-type="issue" data-hovercard-url="/golang/go/issues/42646/hovercard" href="https://github.com/golang/go/issues/42646">golang/go#42646</a>.</p>
<hr>
<p>A full list of all issues fixed can be found in the <a href="https://github.com/golang/go/milestone/184?closed=1">gopls/v0.5.4 milestone</a>. To report a new problem, please file a new issue at <a href="https://golang.org/issues/new" rel="nofollow">https://golang.org/issues/new</a>.</p>
<h2>Thank you to our contributors!</h2>
<p><a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/findleyr/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/findleyr">@findleyr</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/heschik/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/heschik">@heschik</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/pjweinb/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/pjweinb">@pjweinb</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/leitzler/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/leitzler">@leitzler</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/yangwenmai/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/yangwenmai">@yangwenmai</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/matloob/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/matloob">@matloob</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/golopot/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/golopot">@golopot</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/muirdm/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/muirdm">@muirdm</a></p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.4', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.4-pre.1', '<p>gopls/v0.5.4-pre.1</p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.4-pre.1', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.3', '<h1>gopls/v0.5.3</h1>
<h2>Features</h2>
<h3>Automatic updates to <code>go.sum</code></h3>
<p>Previously, <code>go.mod</code>-related quick fixes would not make corresponding changes to your <code>go.sum</code> file. Now, when you add or remove dependencies from your module, your <code>go.sum</code> will be updated accordingly.</p>
<h3>Removed support for <code>go mod tidy</code> on save</h3>
<p>We have removed the support for running <code>go mod tidy</code> on save for <code>go.mod</code> files. It proved to be too slow and expensive to be worth it.</p>
<h2>Experimental</h2>
<h3>Multi-module workspace support</h3>
<p>The proposal described in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="451231859" data-permission-text="Title is private" data-url="https://github.com/golang/go/issues/32394" data-hovercard-type="issue" data-hovercard-url="/golang/go/issues/32394/hovercard" href="https://github.com/golang/go/issues/32394">golang/go#32394</a> is still in development and off by default. See our progress by tracking the multi-module workspace <a href="https://github.com/golang/go/milestone/179">milestone</a> and <a href="https://github.com/golang/go/projects/4#card-47762504">project</a>.</p>
<p>Enable multi-module workspace support by adding the following to your settings:</p>
<div class="highlight highlight-source-js"><pre><span class="pl-s">"gopls"</span>: <span class="pl-kos">{</span>
    <span class="pl-s">"experimentalWorkspaceModule"</span>: <span class="pl-c1">true</span><span class="pl-kos">,</span>
<span class="pl-kos">}</span></pre></div>
<p>With this setting, you will be able to open a directory that contains multiple modules. Most features will work across modules, but some, such as <code>goimports</code>, will not work as expected.</p>
<p>Give this a try if you''re interested in this new feature, but please note that it is still very experimental.</p>
<h3>Fixes</h3>
<p>A list of all issues fixed can be found in the <a href="https://github.com/golang/go/milestone/181?closed=1">gopls/v0.5.3</a> milestone.</p>
<h3>Thank you to our contributors!</h3>
<p><a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/heschik/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/heschik">@heschik</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/findleyr/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/findleyr">@findleyr</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/pjweinb/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/pjweinb">@pjweinb</a></p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.3', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.3-pre.2', '<p>gopls/v0.5.3-pre.2</p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.3-pre.2', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.3-pre.1', '<p>gopls/v0.5.3-pre.1</p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.3-pre.1', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.2', '<h1>gopls/v0.5.2</h1>
<h2>Features</h2>
<p>No new features have been added in this release.</p>
<h2>Experimental</h2>
<p><strong>We have added support for a new <code>allExperiments</code> setting.</strong> By enabling this flag, you will enable all experimental features that we intend to roll out slowly. You can still disable individual settings (<a href="https://github.com/golang/tools/blob/master/gopls/doc/settings.md">full list of settings</a>). In-progress features, such as multi-module workspaces (below), will remain disabled until they are ready for users.</p>
<h3>Improved CPU utilization: <a href="https://github.com/golang/tools/blob/master/gopls/doc/settings.md#experimentaldiagnosticsdelay-timeduration"><code>experimentalDiagnosticsDelay</code></a></h3>
<p><code>experimentalDiagnosticsDelay</code> controls the amount of time that gopls waits after the most recent file modification before computing deep diagnostics. Simple diagnostics (parsing and type-checking) are always run immediately on recently modified packages.</p>
<p>Enable it by setting it to a <a href="https://pkg.go.dev/time#ParseDuration" rel="nofollow">duration string</a>, for example <code>"200ms"</code>. With <code>allExperiments</code>, this is set to <code>"200ms"</code>.</p>
<h3>Improved memory usage for workspaces with multiple folders: <a href="https://github.com/golang/tools/blob/master/gopls/doc/settings.md#experimentalpackagecachekey-bool"><code>experimentalPackageCacheKey</code></a></h3>
<p><code>experimentalPackageCacheKey</code> controls whether to use a coarser cache key for package type information. If you use the gopls <a href="https://github.com/golang/tools/blob/master/gopls/doc/daemon.md">daemon</a>, this may reduce your total memory usage.</p>
<p>Enable it by setting it to <code>true</code>. With <code>allExperiments</code>, this is set to <code>true</code>.</p>
<h3>Multi-module workspace support</h3>
<p>The proposal described in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="451231859" data-permission-text="Title is private" data-url="https://github.com/golang/go/issues/32394" data-hovercard-type="issue" data-hovercard-url="/golang/go/issues/32394/hovercard" href="https://github.com/golang/go/issues/32394">golang/go#32394</a> is still in development and off by default. See our progress by tracking the multi-module workspace <a href="https://github.com/golang/go/milestone/179">milestone</a> and <a href="https://github.com/golang/go/projects/4#card-47762504">project</a>.</p>
<p>Enable multi-module workspace support by adding the following to your settings:</p>
<div class="highlight highlight-source-js"><pre><span class="pl-s">"gopls"</span>: <span class="pl-kos">{</span>
    <span class="pl-s">"experimentalWorkspaceModule"</span>: <span class="pl-c1">true</span><span class="pl-kos">,</span>
<span class="pl-kos">}</span></pre></div>
<p>With this setting, you will be able to open a directory that contains multiple modules. Most features will work across modules, but some, such as <code>goimports</code>, will not work as expected.</p>
<p>Give this a try if you''re interested in this new feature, but please note that it is still very experimental.</p>
<h3>Support for <a href="https://microsoft.github.io/language-server-protocol/specifications/specification-3-16/#textDocument_semanticTokens" rel="nofollow">semantic tokens</a></h3>
<p>This is a new, unreleased LSP feature that provides additional syntax highlighting. In advance of this new LSP version, we have added preliminary support for this feature. Enable it by setting:</p>
<div class="highlight highlight-source-js"><pre><span class="pl-s">"gopls"</span>: <span class="pl-kos">{</span>
    <span class="pl-s">"semanticTokens"</span>: <span class="pl-c1">true</span><span class="pl-kos">,</span>
<span class="pl-kos">}</span></pre></div>
<p>It will not be enabled with <code>allExperiments</code>.</p>
<h3>Fixes</h3>
<p>A list of all issues fixed can be found in the <a href="https://github.com/golang/go/milestone/174?closed=1">gopls/v0.5.2</a> milestone.</p>
<h3>For editor clients</h3>
<p>All command names have been given <code>gopls.</code> prefixes, to avoid conflicting with commands registered by other language servers.<br>
This should not have affected any clients.</p>
<h3>Thank you to our contributors!</h3>
<p><a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/heschik/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/heschik">@heschik</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/findleyr/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/findleyr">@findleyr</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/dandua98/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/dandua98">@dandua98</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/pjweinb/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/pjweinb">@pjweinb</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/leitzler/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/leitzler">@leitzler</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/kortschak/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/kortschak">@kortschak</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/bcmills/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/bcmills">@bcmills</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/hyangah/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/hyangah">@hyangah</a> <a class="user-mention" data-hovercard-type="user" data-hovercard-url="/users/jadekler/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/jadekler">@jadekler</a></p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.2', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.5.2-pre.2', '<p>gopls/v0.5.2-pre.2</p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.5.2-pre.2', NULL);
INSERT INTO public.releases (title, content, link, id) VALUES ('gopls/v0.6.0-pre.1', '<p>gopls/v0.6.0-pre.1</p>', 'https://github.com/golang/tools/releases/tag/gopls%2Fv0.6.0-pre.1', NULL);
