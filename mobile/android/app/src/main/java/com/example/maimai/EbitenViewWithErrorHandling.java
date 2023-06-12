package com.example.maimai;

import android.content.Context;
import android.util.AttributeSet;

import gixel.maimai.EbitenView;

class EbitenViewWithErrorHandling extends EbitenView {
    public EbitenViewWithErrorHandling(Context context) {
        super(context);
        System.out.println("RON");
    }

    public EbitenViewWithErrorHandling(Context context, AttributeSet attributeSet) {
        super(context, attributeSet);
    }

    @Override
    protected void onErrorOnGameUpdate(Exception e) {
        // You can define your own error handling e.g., using Crashlytics.
        // e.g., Crashlytics.logException(e);
        System.out.println(e);
        super.onErrorOnGameUpdate(e);
    }
}