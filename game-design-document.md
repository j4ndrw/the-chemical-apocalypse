# The Chemical Apocalypse Design Document

## Game Overview

**Title**: The Chemical Apocalypse
**Genre**: Hybrid (Roguelike, Turn-Based Mechanics)

The Chemical Apocalypse is an innovative game that combines elements of roguelikes with a unique turn-based combat system. Players will navigate a world filled with chemical-based abilities, engaging in strategic battles against enemies that are themselves amalgams of various chemicals.

## Core Mechanics

### 1. Item Discovery and Synergy
- Players will explore the game world to discover a variety of items.
- Items can be combined to create synergies, enhancing the player's abilities and combat effectiveness.

### 2. Combat System
- The combat system features a semi-turn-based mode when engaging with one or more enemies:
  - Players and enemies alternate turns, allowing for strategic planning.
  - A **degree of violence** system enables multiple actions to be performed without a strict limit.

### 3. Action Limitations and Negative Properties
- While there is no limit to the number of actions a player can perform, each action contributes to the overall degree of violence of the final reaction:
  - The more actions performed, the more violent the final reaction will be.
  - To balance this mechanic, players may incur **negative properties** that amplify the negative effects of their actions as they chain more reactions.
  - Healing actions will also be limited to ensure that players cannot easily negate the consequences of their actions. This can be implemented by providing a set number of healing opportunities per floor or through other resource constraints.
  - Alternatively, players can discover or equip **stabilizers** that allow them to perform additional actions without incurring the usual penalties, providing strategic options for managing the degree of violence.

### 4. Real-Time Reactions
- Defensive reactions occur in real-time, despite the turn-based nature of actions:
  - Players can block or dodge enemy attacks by timing their responses correctly, potentially reducing or avoiding damage.
  - This mechanic encourages quick decision-making and enhances the overall combat experience.

### 5. Dynamic Action Types
- Players and enemies will have access to three distinct types of actions:
  - **Tactical Actions**: Turn-based attacks that cannot be executed reactively.
  - **Reactive Actions**: Defensive maneuvers (e.g., dodging) that can be performed in response to enemy actions.
  - **Preparable Actions**: Actions that can be executed during the opponent's turn, including healing, buffing, or inflicting debuffs.
- Each action will be mapped to a specific key on the keyboard or controller, allowing for customizable controls and efficient gameplay.

### 6. Action Classification
- Actions will be categorized as follows:
  - **Reactive**: Executed in response to an enemy action (e.g., dodging).
  - **Tactical**: Executed during the player's turn (e.g., attacking).
  - **Preparable**: Executed during the opponent's turn (e.g., healing or buffing).

### 7. Enemy Design
- Each enemy will be an amalgam of various chemicals, exhibiting predictable actions based on their chemical composition.
- This design choice facilitates fast-paced and satisfying gameplay, particularly in the later stages of the game.
- Players will be able to anticipate enemy behaviors, allowing for effective chaining of actions against multiple foes, enhancing the excitement of combat.

## Game Premise

In The Chemical Apocalypse, players begin with a fundamental set of chemical abilities that provide the basis for combat:
- A basic tactical attack
- A basic reactive defense
- A basic preparable stat-changing action (e.g., healing, buffing, debuffing)

### Degree of Violence and Reaction Spectrum
- Each action contributes to a spectrum that ranges from **0 to 100**:
  - **0** represents an inert reaction (no violence).
  - **100** represents a violent reaction (maximum violence).
- The overall impact of a reaction is determined by the aggregated points on this spectrum. The higher the value, the more pronounced the effects of the action, including any associated drawbacks.

This duality of reactions encourages players to strategize their actions based on potential risks and rewards, fostering a dynamic and engaging gameplay experience.

## Additional Notes
- Consider developing a crafting system in the future. This system could allow players to combine various items and chemicals to create new abilities or enhance existing ones, adding depth and complexity to gameplay as the game matures.
