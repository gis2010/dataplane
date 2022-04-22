import { createState, useState as useHookState } from '@hookstate/core';

// Global run state - run_id gets added to global state when pipeline is run
export const globalDeploymentState = createState({
    selectedRunID: null,
    runIDs: null,
    runTrigger: 0,
    tableRunTrigger: 0,
    onLoadTrigger: 0,
    onChangeTrigger: 0,
    node_id: null,
    isRunning: false,
});

export const useGlobalDeploymentState = () => useHookState(globalDeploymentState);