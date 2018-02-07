/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */
package org.apache.plc4x.java.ads.api.commands;

import org.apache.plc4x.java.ads.api.commands.types.NotificationHandle;
import org.apache.plc4x.java.ads.api.generic.ADSData;
import org.apache.plc4x.java.ads.api.generic.AMSHeader;
import org.apache.plc4x.java.ads.api.generic.AMSTCPHeader;
import org.apache.plc4x.java.ads.api.generic.AMSTCPPaket;
import org.apache.plc4x.java.ads.api.generic.types.*;

/**
 * One before defined notification is deleted in an ADS device.
 */
public class ADSDeleteDeviceNotificationRequest extends AMSTCPPaket {

    /**
     * 4 bytes	Handle of notification
     */
    private final NotificationHandle notificationHandle;

    public ADSDeleteDeviceNotificationRequest(AMSTCPHeader amstcpHeader, AMSHeader amsHeader, NotificationHandle notificationHandle) {
        super(amstcpHeader, amsHeader);
        this.notificationHandle = notificationHandle;
    }

    public ADSDeleteDeviceNotificationRequest(AMSHeader amsHeader, NotificationHandle notificationHandle) {
        super(amsHeader);
        this.notificationHandle = notificationHandle;
    }

    public ADSDeleteDeviceNotificationRequest(AMSNetId targetAmsNetId, AMSPort targetAmsPort, AMSNetId sourceAmsNetId, AMSPort sourceAmsPort, Invoke invokeId, Data nData, NotificationHandle notificationHandle) {
        super(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, invokeId, nData);
        this.notificationHandle = notificationHandle;
    }

    @Override
    public ADSData getAdsData() {
        return buildADSData(notificationHandle);
    }

    @Override
    public Command getCommandId() {
        return Command.ADS_Delete_Device_Notification;
    }

    @Override
    public State getStateId() {
        return State.ADS_REQUEST_TCP;
    }
}
