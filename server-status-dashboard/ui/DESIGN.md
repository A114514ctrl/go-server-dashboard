---
name: Obsidian Terminal
colors:
  surface: '#131313'
  surface-dim: '#131313'
  surface-bright: '#393939'
  surface-container-lowest: '#0e0e0e'
  surface-container-low: '#1c1b1b'
  surface-container: '#201f1f'
  surface-container-high: '#2a2a2a'
  surface-container-highest: '#353534'
  on-surface: '#e5e2e1'
  on-surface-variant: '#b9ccb2'
  inverse-surface: '#e5e2e1'
  inverse-on-surface: '#313030'
  outline: '#84967e'
  outline-variant: '#3b4b37'
  surface-tint: '#00e639'
  primary: '#ebffe2'
  on-primary: '#003907'
  primary-container: '#00ff41'
  on-primary-container: '#007117'
  inverse-primary: '#006e16'
  secondary: '#dcfdff'
  on-secondary: '#00373a'
  secondary-container: '#00f1fd'
  on-secondary-container: '#006a6f'
  tertiary: '#fff7f6'
  on-tertiary: '#690006'
  tertiary-container: '#ffd2cd'
  on-tertiary-container: '#c40015'
  error: '#ffb4ab'
  on-error: '#690005'
  error-container: '#93000a'
  on-error-container: '#ffdad6'
  primary-fixed: '#72ff70'
  primary-fixed-dim: '#00e639'
  on-primary-fixed: '#002203'
  on-primary-fixed-variant: '#00530e'
  secondary-fixed: '#6ff6ff'
  secondary-fixed-dim: '#00dce6'
  on-secondary-fixed: '#002022'
  on-secondary-fixed-variant: '#004f53'
  tertiary-fixed: '#ffdad6'
  tertiary-fixed-dim: '#ffb4ab'
  on-tertiary-fixed: '#410002'
  on-tertiary-fixed-variant: '#93000c'
  background: '#131313'
  on-background: '#e5e2e1'
  surface-variant: '#353534'
typography:
  display-lg:
    fontFamily: JetBrains Mono
    fontSize: 48px
    fontWeight: '700'
    lineHeight: 56px
    letterSpacing: -0.02em
  headline-lg:
    fontFamily: JetBrains Mono
    fontSize: 32px
    fontWeight: '600'
    lineHeight: 40px
    letterSpacing: -0.01em
  headline-md:
    fontFamily: JetBrains Mono
    fontSize: 24px
    fontWeight: '600'
    lineHeight: 32px
  body-lg:
    fontFamily: JetBrains Mono
    fontSize: 16px
    fontWeight: '400'
    lineHeight: 24px
  body-md:
    fontFamily: JetBrains Mono
    fontSize: 14px
    fontWeight: '400'
    lineHeight: 20px
  label-lg:
    fontFamily: JetBrains Mono
    fontSize: 12px
    fontWeight: '500'
    lineHeight: 16px
    letterSpacing: 0.1em
  label-sm:
    fontFamily: JetBrains Mono
    fontSize: 10px
    fontWeight: '500'
    lineHeight: 14px
    letterSpacing: 0.05em
  mono-data:
    fontFamily: JetBrains Mono
    fontSize: 13px
    fontWeight: '400'
    lineHeight: 18px
spacing:
  unit: 4px
  gutter: 16px
  margin-mobile: 16px
  margin-desktop: 32px
  container-max: 1440px
---

## Brand & Style
The design system is an ultra-dark, high-performance interface designed for systems engineers and developers who thrive in low-light, high-density data environments. It evokes the feeling of a mission-critical terminal—efficient, precise, and authoritative.

The aesthetic blends **Minimalism** with **Retro-Futurism**. It prioritizes function over form, utilizing a strict grid and thin-line borders. Visual interest is generated through purposeful "glow" states (neon luminescence) against a pure black void, ensuring that the most critical system statuses are immediately perceptible through chromatic cues.

## Colors
This design system utilizes a "Void-Black" foundation to maximize contrast and reduce eye strain during long monitoring sessions. 

- **Matrix Green (#00FF41):** Used for "System Healthy" states, successful pings, and primary action buttons.
- **Cyber Cyan (#00F3FF):** Used for informational metrics, active processes, and secondary data visualizations.
- **Alert Red (#FF3131):** Reserved strictly for critical failures, high CPU load, and "Down" status.
- **Surface & Borders:** Grays are kept extremely dark (#080808) to maintain the depth of the black background, with thin borders (#1A1A1A) defining the structural skeleton.

## Typography
The typography is strictly monospaced to reinforce the technical nature of the dashboard. JetBrains Mono is used for its exceptional legibility in dense data sets.

- **Headlines:** Use Bold weights with slight negative letter spacing to create a compact, modern feel.
- **Labels:** Small labels use uppercase with tracking (letter-spacing) to mimic the look of legacy hardware interfaces.
- **Data Points:** Real-time metrics and terminal logs use `mono-data` for perfect vertical alignment of characters, facilitating quick scanning of IP addresses and timestamps.

## Layout & Spacing
The layout follows a **Fixed Grid** philosophy, utilizing a 12-column system for desktop. The UI is structured into "Modules" that snap to the grid.

- **Grid Lines:** On large screens, subtle 1px vertical and horizontal lines can be visible in the background to emphasize the engineering precision.
- **Density:** High information density is preferred. Padding is minimal (8px or 12px within cards) to maximize the amount of visible data on a single screen.
- **Adaptation:** On mobile, the grid collapses to a single column. Horizontal scrolling is permitted for wide data tables to preserve the monospace alignment of columns.

## Elevation & Depth
In a pure black environment, traditional shadows are ineffective. Instead, depth is communicated through **Luminance and Outlines**.

- **Level 0 (Background):** Pure black (#000000).
- **Level 1 (Cards/Containers):** Subtle 1px border (#1A1A1A). No fill, or a very dark fill (#080808).
- **Interactive State (Glow):** When an element is active or hovered, it gains a subtle outer glow (2px to 8px blur) using the component's accent color (Green or Cyan) at 30% opacity.
- **Layering:** Modals or overlays use a solid 1px Matrix Green border to "cut" through the background, rather than using drop shadows.

## Shapes
The shape language is strictly **Sharp (0px)**. Rounded corners are avoided to maintain a brutalist, high-tech aesthetic.

- **Buttons & Inputs:** Rectangular with hard 90-degree angles.
- **Data Visualization:** Bars in charts should be flat-topped.
- **Dividers:** Horizontal and vertical lines are exactly 1px thick.

## Components

### Buttons
- **Primary:** 1px Matrix Green border, Green text, transparent background. On hover: Matrix Green background with Black text and a subtle green glow.
- **Secondary:** 1px Cyber Cyan border, Cyan text. 

### Data-Heavy Cards
- **Header:** 1px bottom border separating the title from the content. Title in `label-lg` style.
- **Body:** Contains condensed metrics. Values in Green or Cyan to distinguish them from the gray labels.

### Circular Progress Meters
- **Visuals:** Non-rounded segments. Background track is #1A1A1A. Active track uses a "Neon" effect (primary color with a thin centered line of a lighter shade).
- **Status:** Color shifts from Green (0-70%) to Cyan (71-90%) to Red (91-100%).

### Terminal-Style Lists
- **Rows:** Alternating subtle background tints (#050505) for readability.
- **Prefixes:** Every line begins with a prompt character (`>`) or a timestamp in gray.
- **Status Indicators:** Use a simple square glyph (■) instead of a circle for the status dot.

### Input Fields
- **Style:** 1px border on the bottom only, or a full 1px border. Background is transparent.
- **Focus:** The border changes to Matrix Green with a blinking vertical cursor (pipe character `|`).