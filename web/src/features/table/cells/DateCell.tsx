import React from 'react';

import { CalendarLtr20Regular as DateIcon, Clock20Regular as TimeIcon } from '@fluentui/react-icons';
import classNames from 'classnames';
import { format } from 'date-fns';

import styles from './cell.module.scss';

interface DateCellProps {
  value: string | Date;
  className?: string;
}

function DateCell(props: DateCellProps) {
  let date = props.value;
  if (typeof props.value === 'string') {
    date = new Date(props.value);
  }
  let displayDate = '';
  let displayTime = '';
  if (date != null && date.valueOf() !== 0) {
    displayDate = format(date as Date, 'yyyy-MM-dd');
    displayTime = format(date as Date, 'HH:mm:ss');
  } else {
    displayDate = '';
  }

  return (
    <div className={classNames(styles.cell, styles.alignLeft, styles.booleanCell, styles.DateCell, props.className)}>
      <div className={styles.dateRow}>
        <DateIcon />
        {displayDate}
      </div>
      <div className={styles.timeRow}>
        <TimeIcon />
        {displayTime}
      </div>
    </div>
  );
}

const DateCellMemo = React.memo(DateCell);
export default DateCellMemo;
