import React, { useContext } from 'react';
import { WebSocketContext } from '../WebSocket';
import { DetailsList, IDetailsListProps } from '@fluentui/react';
import { IControlProps } from './IControlProps'


export const MyDetailsList = React.memo<IControlProps>(({control, parentDisabled}) => {

    const ws = useContext(WebSocketContext);

    const detailsListProps: IDetailsListProps = {
        items: control.items,
        columns: control.columns,
        layoutMode: control.layoutMode
    };

    return <DetailsList {...detailsListProps} />
})