<template>
    <div class="control-general">
        <span v-b-tooltip.hover
              :title="'Immediately stop all robots (' + Object.keys(keymapHalt)[0] + ')'">
            <b-button v-hotkey="keymapHalt"
                      ref="btnHalt"
                      v-on:click="send('halt')"
                      v-bind:disabled="halted">
                Halt
            </b-button>
        </span>
        <span v-b-tooltip.hover
              :title="'Robots have to keep distance to the ball (' + Object.keys(keymapStop)[0] + ')'">
            <b-button v-hotkey="keymapStop"
                      ref="btnStop"
                      v-on:click="send('stop')"
                      v-bind:disabled="stopped || !stopAllowed">
                Stop
            </b-button>
        </span>
        <span v-b-tooltip.hover
              :title="'Restart the game in draw situations (' + Object.keys(keymapForceStart)[0] + ')'">
            <b-button v-hotkey="keymapForceStart"
                      ref="btnForceStart"
                      v-on:click="send('forceStart')"
                      v-bind:disabled="!stopped || !forceStartAllowed">
                Force Start
            </b-button>
        </span>
        <span v-b-tooltip.hover
              :title="'Continue game after a prepare state (' + Object.keys(keymapNormalStart)[0] + ')'">
            <b-button v-hotkey="keymapNormalStart"
                      ref="btnNormalStart"
                      v-on:click="send('normalStart')"
                      v-bind:disabled="!prepareSth || !normalStartAllowed">
                Normal Start
            </b-button>
        </span>
        <span v-b-tooltip.hover
              :title="'Continue based on last game event (' + Object.keys(keymapContinue)[0] + ')'">
            <b-button v-hotkey="keymapContinue"
                      ref="btnContinue"
                      v-on:click="triggerContinue"
                      v-bind:disabled="!continuePossible">
                Continue
            </b-button>
        </span>
        <div class="btn-group-toggle btn-group" v-hotkey="keymapToggleAutoRef">
            <label v-b-tooltip.hover
                   :title="'Enable automatic continuation based on game events (' + Object.keys(keymapToggleAutoRef)[0] + ')'"
                   :class="{btn:true, 'btn-secondary': true, active: autoContinue}"
                   @click="setAutoContinue(true)">
                Auto
            </label>
            <label v-b-tooltip.hover
                   :title="'Disable automatic continuation based on game events (' + Object.keys(keymapToggleAutoRef)[0] + ')'"
                   :class="{btn:true, 'btn-secondary': true, active: !autoContinue}"
                   @click="setAutoContinue(false)">
                Manual
            </label>
        </div>
    </div>
</template>

<script>
    import {isNonPausedStage, isPreStage} from "../../refereeState";

    export default {
        name: "ControlGeneral",
        methods: {
            send: function (command) {
                this.$socket.sendObj({command: {commandType: command}})
            },
            triggerContinue() {
                this.$socket.sendObj({trigger: {triggerType: 'continue'}})
            },
            setAutoContinue(enabled) {
                this.$socket.sendObj({
                    'modify': {'autoContinue': enabled}
                })
            },
        },
        computed: {
            keymapHalt() {
                return {
                    'esc': () => {
                        if (!this.$refs.btnHalt.disabled) {
                            this.send('halt')
                        }
                    }
                }
            },
            keymapStop() {
                return {
                    'ctrl+alt+numpad 0': () => {
                        if (!this.$refs.btnStop.disabled) {
                            this.send('stop')
                        }
                    }
                }
            },
            keymapForceStart() {
                return {
                    'ctrl+alt+numpad 5': () => {
                        if (!this.$refs.btnForceStart.disabled) {
                            this.send('forceStart')
                        }
                    }
                }
            },
            keymapNormalStart() {
                return {
                    'ctrl+alt+numpad 8': () => {
                        if (!this.$refs.btnNormalStart.disabled) {
                            this.send('normalStart')
                        }
                    }
                }
            },
            keymapContinue() {
                return {
                    'ctrl+alt+space': () => {
                        if (!this.$refs.btnContinue.disabled) {
                            this.triggerContinue()
                        }
                    }
                }
            },
            keymapToggleAutoRef() {
                return {
                    'ctrl+alt+numpad 2': () => {
                        this.setAutoContinue(!this.autoContinue)
                    }
                }
            },
            state() {
                return this.$store.state.refBoxState
            },
            halted() {
                return this.state.command === 'halt';
            },
            stopped() {
                return this.state.command === 'stop';
            },
            prepareSth() {
                return this.state.command === 'kickoff' || this.state.command === 'penalty';
            },
            forceStartAllowed() {
                return isNonPausedStage(this.state);
            },
            normalStartAllowed() {
                return isNonPausedStage(this.state) || this.state.command === 'kickoff';
            },
            stopAllowed() {
                return isNonPausedStage(this.$store.state.refBoxState)
                    || isPreStage(this.$store.state.refBoxState);
            },
            autoContinue() {
                return this.state.autoContinue;
            },
            continuePossible() {
                return this.state.nextCommand !== '';
            }
        }
    }
</script>

<style scoped>
    button, .btn-group {
        margin-right: 0.5em;
        margin-bottom: 0.5em;
    }
</style>