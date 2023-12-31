// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {combineReducers} from 'redux';

import type {AdminRequestsStatuses, RequestStatusType} from '@mattermost/types/requests';

import {AdminTypes} from 'mattermost-redux/action_types';
import type {GenericAction} from 'mattermost-redux/types/actions';

import {handleRequest, initialRequestState} from './helpers';

function createCompliance(state: RequestStatusType = initialRequestState(), action: GenericAction): RequestStatusType {
    return handleRequest(
        AdminTypes.CREATE_COMPLIANCE_REQUEST,
        AdminTypes.CREATE_COMPLIANCE_SUCCESS,
        AdminTypes.CREATE_COMPLIANCE_FAILURE,
        state,
        action,
    );
}

export default (combineReducers({
    createCompliance,
}) as (b: AdminRequestsStatuses, a: GenericAction) => AdminRequestsStatuses);
