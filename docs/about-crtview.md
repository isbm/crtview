Reasons why [cview](https://gitlab.com/tslocum/cview) was forked.

# Fork

As [rivo](https://github.com/rivo), the creator and sole maintainer of
tview, explains his reviewing and merging process in a
[GitHub comment](https://github.com/rivo/tview/pull/298#issuecomment-559373851),
he does not have the necessary time or interest to review, discuss and
merge pull requests. For more details, follow the link
above. Nevertheless, this gave a number of reasons to fork it as
[cview](https://gitlab.com/tslocum/cview) project, which supposed to
solve **tview**-related issues.

Problem, however, that [cview](https://gitlab.com/tslocum/cview)
introduces a number of differences, mainly completely removing chained
calls, which results into unnecessarily overly bloated code for your app.

# Differences from tview

API of the widgets and primitives are simply more rich and constantly
growing, on demand. This document is not intended to list them all.

The **crtview** inherits all the differences between **tview** and
**cview**.

Chained calls are aimed to be restored back and improved on their own.
