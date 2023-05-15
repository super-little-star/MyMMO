using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UILoading : UIBase
{
    public Slider S_Progress;

    public void SetSliderValue(float value)
    {
        S_Progress.value = value;
    }
}
