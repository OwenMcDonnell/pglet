import React, { useContext } from 'react';
import { WebSocketContext } from '../WebSocket';
import { shallowEqual, useSelector } from 'react-redux'
import { DetailsList, IDetailsListProps, IColumn } from '@fluentui/react';
import { IControlProps } from './IControlProps'


export const MyDetailsList = React.memo<IControlProps>(({control}) => {
    
  const ws = useContext(WebSocketContext);

    const detailsListProps: IDetailsListProps = {
        items: [],
        columns: [],
        layoutMode: control.layoutMode
    };

    detailsListProps.items = useSelector<any, any[]>((state: any) => control.c.map((childId: any) =>
    state.page.controls[childId])
      .filter((lc: any) => lc.t === 'listItem')
      .map((lc: any) => ({ key: lc.key, value: lc.value, type: lc.type})), shallowEqual);  


    detailsListProps.columns = useSelector<any, any[]>((state: any) => control.c.map((childId: any) =>
    state.page.controls[childId])
      .filter((lc: any) => lc.t === 'listColumn')
      .map((lc: any) => ({ key: lc.key, name: lc.value, minWidth: 80})), shallowEqual);  

    console.log("detailsListProps");
    console.log(detailsListProps);

    return <DetailsList {...detailsListProps} />
})